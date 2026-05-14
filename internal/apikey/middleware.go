package apikey

import (
	"net/http"
)

type Middleware interface {
	Protect(next http.Handler) http.Handler
}

type middleware struct {
	apiKey string
}

func NewMiddleware(apiKey string) Middleware {
	return &middleware{
		apiKey: apiKey,
	}
}

func (m *middleware) Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if apiKey == "" {
			http.Error(w, "missing api key", http.StatusUnauthorized)
			return
		}

		if apiKey != m.apiKey {
			http.Error(w, "invalid api key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
