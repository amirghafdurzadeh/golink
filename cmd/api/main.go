package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/amirghafdurzadeh/golink/internal/apikey"
	"github.com/amirghafdurzadeh/golink/internal/config"
	"github.com/amirghafdurzadeh/golink/internal/health"
	"github.com/amirghafdurzadeh/golink/internal/link"
	"github.com/amirghafdurzadeh/golink/internal/redirect"
	"github.com/amirghafdurzadeh/golink/internal/router"
)

func main() {
	// config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// infrastructure
	startupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postgresPool := mustPostgres(startupCtx, cfg)
	defer postgresPool.Close()

	redisClient := mustRedis(startupCtx, cfg)
	defer redisClient.Close()

	// repositories
	linkRepository := link.NewPostgresRepository(postgresPool)

	// caches
	linkCache := link.NewRedisCache(redisClient, 24*time.Hour)

	// services
	linkService := link.NewService(linkRepository, linkCache, cfg.ShortCodeLength)

	// handlers
	healthHandler := health.NewHandler(postgresPool)
	redirectHandler := redirect.NewHandler(postgresPool)
	linkHandler := link.NewHandler(linkService)

	// middleware
	apiKeyMiddleware := apikey.NewMiddleware(cfg.APIKey)

	// routers
	rootMux := http.NewServeMux()
	router.RegisterRoutes(
		rootMux,
		linkHandler,
		redirectHandler,
		healthHandler,
		apiKeyMiddleware,
	)

	// server
	server := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           rootMux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("server starting on port %s", cfg.AppPort)

	if err := server.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
