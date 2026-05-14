package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/amirghafdurzadeh/golink/internal/apikey"
	"github.com/amirghafdurzadeh/golink/internal/config"
	"github.com/amirghafdurzadeh/golink/internal/database"
	"github.com/amirghafdurzadeh/golink/internal/health"
	"github.com/amirghafdurzadeh/golink/internal/link"
	"github.com/amirghafdurzadeh/golink/internal/redirect"
)

func main() {
	// config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// infrastructure
	startupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postgresPool, err := database.NewPostgres(startupCtx, cfg.PostgresURL)
	if err != nil {
		log.Fatal(err)
	}
	defer postgresPool.Close()

	redisClient, err := database.NewRedis(startupCtx, cfg.RedisAddr, cfg.RedisPassword)
	if err != nil {
		log.Fatal(err)
	}

	// repositories
	linkRepository := link.NewPostgresRepository(postgresPool)

	// cashes
	linkCache := link.NewRedisCache(redisClient, 24*time.Hour)

	// services
	linkService := link.NewService(linkRepository, linkCache, cfg.ShortCodeLength)

	// handlers
	healthHandler := health.NewHandler(postgresPool)
	redirectHandler := redirect.NewHandler(postgresPool)
	linkHandler := link.NewHandler(linkService)

	// middleware
	apiKeyMiddleware := apikey.NewMiddleware(
		cfg.APIKey,
	)

	// routers
	apiV1Mux := http.NewServeMux()
	{
		apiV1Mux.HandleFunc("POST /links", linkHandler.Create)
		apiV1Mux.HandleFunc("GET /links/{code}", linkHandler.Get)
		apiV1Mux.HandleFunc("DELETE /links/{code}", linkHandler.Delete)
	}

	rootMux := http.NewServeMux()
	{
		rootMux.HandleFunc("GET /health", healthHandler.Health)
		rootMux.HandleFunc("GET /r/{code}", redirectHandler.Redirect)

		rootMux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiKeyMiddleware.Protect(apiV1Mux)))
	}

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
