package middleware

import (
	"net/http"
)

type APIKeyMiddleware struct {
	apiKey string
}

func NewAPIKeyMiddleware(apiKey string) *APIKeyMiddleware {
	return &APIKeyMiddleware{
		apiKey: apiKey,
	}
}

func (m *APIKeyMiddleware) Protect(next http.Handler) http.Handler {
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
