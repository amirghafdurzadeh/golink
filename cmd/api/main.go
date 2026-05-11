package main

import (
	"log"
	"net/http"
	"os"

	"github.com/amirghafdurzadeh/golink/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	apiV1Mux := http.NewServeMux()
	{
		apiV1Mux.HandleFunc("POST /links", handler.CreateLink)
		apiV1Mux.HandleFunc("GET /links/{code}", handler.GetLink)
		apiV1Mux.HandleFunc("DELETE /links/{code}", handler.DeleteLink)
		apiV1Mux.HandleFunc("GET /links/{code}/stats", handler.GetLinkStats)
	}

	rootMux := http.NewServeMux()
	{
		rootMux.HandleFunc("GET /health", handler.Health)
		rootMux.HandleFunc("GET /r/{code}", handler.Redirect)

		rootMux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1Mux))
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, rootMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
