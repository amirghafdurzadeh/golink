package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/amirghafdurzadeh/golink/internal/config"
	"github.com/amirghafdurzadeh/golink/internal/db"
	"github.com/amirghafdurzadeh/golink/internal/handler"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database, err := db.New(ctx, cfg.PostgresConnURL, cfg.RedisAddr, cfg.RedisPassword)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	h := handler.New(database)

	apiV1Mux := http.NewServeMux()
	{
		apiV1Mux.HandleFunc("POST /links", h.CreateLink)
		apiV1Mux.HandleFunc("GET /links/{code}", h.GetLink)
		apiV1Mux.HandleFunc("DELETE /links/{code}", h.DeleteLink)
		apiV1Mux.HandleFunc("GET /links/{code}/stats", h.GetLinkStats)
	}

	rootMux := http.NewServeMux()
	{
		rootMux.HandleFunc("GET /health", h.Health)
		rootMux.HandleFunc("GET /r/{code}", h.Redirect)

		rootMux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1Mux))
	}

	log.Printf("Server starting on port %s", cfg.AppPort)
	if err := http.ListenAndServe(":"+cfg.AppPort, rootMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
