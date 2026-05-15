package apikey

import (
	"net/http"

	"github.com/amirghafdurzadeh/golink/internal/apikey"
)

type Middleware interface {
	Protect(next http.Handler) http.Handler
}

type middleware struct {
	service apikey.Service
}

func NewMiddleware(service apikey.Service) Middleware {
	return &middleware{
		service: service,
	}
}

func (m *middleware) Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if err := m.service.Validate(apiKey); err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
