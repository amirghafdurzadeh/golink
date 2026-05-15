package router

import (
	"net/http"

	"github.com/amirghafdurzadeh/golink/internal/apikey"
	"github.com/amirghafdurzadeh/golink/internal/health"
	"github.com/amirghafdurzadeh/golink/internal/link"
	"github.com/amirghafdurzadeh/golink/internal/redirect"
)

func RegisterRoutes(
	mux *http.ServeMux,
	linkHandler link.Handler,
	redirectHandler redirect.Handler,
	healthHandler health.Handler,
	apiKeyMiddleware apikey.Middleware,
) {
	mux.HandleFunc("GET /health", healthHandler.Health)
	mux.HandleFunc("GET /r/{code}", redirectHandler.Redirect)

	apiV1 := http.NewServeMux()
	RegisterAPIV1Routes(apiV1, linkHandler)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiKeyMiddleware.Protect(apiV1)))
}

func RegisterAPIV1Routes(
	mux *http.ServeMux,
	linkHandler link.Handler,
) {
	mux.HandleFunc("POST /links", linkHandler.Create)
	mux.HandleFunc("GET /links/{code}", linkHandler.Get)
	mux.HandleFunc("DELETE /links/{code}", linkHandler.Delete)
}
