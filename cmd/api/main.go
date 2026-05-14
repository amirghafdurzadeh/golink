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
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.NewPostgres(ctx, cfg.PostgresConnURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	redisClient, err := database.NewRedis(ctx, cfg.RedisAddr, cfg.RedisPassword)
	if err != nil {
		log.Fatal(err)
	}

	// repositories
	linkRepository := link.NewRepository(db)

	// services
	linkService := link.NewService(linkRepository)

	// handlers
	healthHandler := health.NewHandler(db)
	redirectHandler := redirect.NewHandler(db)
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

	log.Printf("Server starting on port %s", cfg.AppPort)
	if err := http.ListenAndServe(":"+cfg.AppPort, rootMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
