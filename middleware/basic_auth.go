package middleware

import (
	"net/http"
)

// BasicAuthMiddleware validates username and password from the request header
func BasicAuthMiddleware(next http.Handler, validUser, validPass string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != validUser || password != validPass {
			http.Error(w, "Unauthorized - Invalid credentials", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
