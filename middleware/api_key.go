package middleware

import (
	"net/http"
)

// APIKeyMiddleware validates API key from the request header
func APIKeyMiddleware(next http.Handler, validKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientKey := r.Header.Get("X-API-Key")
		if clientKey != validKey {
			http.Error(w, "Unauthorized - Invalid API Key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
