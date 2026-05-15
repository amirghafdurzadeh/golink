package http

import (
	"net/http"

	"github.com/amirghafdurzadeh/golink/internal/app"
	"github.com/amirghafdurzadeh/golink/internal/redirect"
	"github.com/amirghafdurzadeh/golink/internal/transport/http/apikey"
	"github.com/amirghafdurzadeh/golink/internal/transport/http/health"
	"github.com/amirghafdurzadeh/golink/internal/transport/http/link"
)

func newRouter(services app.Services) http.Handler {
	// handlers
	healthHandler := health.NewHandler(services.Health())
	redirectHandler := redirect.NewHandler()
	linkHandler := link.NewHandler(services.Link())

	// middleware
	apiKeyMiddleware := apikey.NewMiddleware(services.APIKey())

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler.Health)
	mux.HandleFunc("GET /r/{code}", redirectHandler.Redirect)

	apiV1 := http.NewServeMux()
	apiV1.HandleFunc("POST /links", linkHandler.Create)
	apiV1.HandleFunc("GET /links/{code}", linkHandler.Get)
	apiV1.HandleFunc("DELETE /links/{code}", linkHandler.Delete)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiKeyMiddleware.Protect(apiV1)))

	return mux
}
