package main

import (
	"fmt"
	"go-auth-example/middleware"
	"log"
	"net/http"
)

const (
	APIKey        = "APIKEY-SECRET"
	BasicAuthUser = "user"
	BasicAuthPass = "password"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! You've successfully accessed a protected route.")
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! This is a public route.")
}

func main() {
	// Handlers for different authentication methods
	apiKeyHandler := middleware.APIKeyMiddleware(http.HandlerFunc(mainHandler), APIKey)
	basicAuthHandler := middleware.BasicAuthMiddleware(http.HandlerFunc(mainHandler), BasicAuthUser, BasicAuthPass)
	jwtHandler := middleware.JWTMiddleware(http.HandlerFunc(mainHandler))
	pubHandler := http.HandlerFunc(publicHandler)

	// Register routes with different authentication middleware
	http.Handle("/api-key", apiKeyHandler)
	http.Handle("/basic-auth", basicAuthHandler)
	http.Handle("/jwt", jwtHandler)
	http.Handle("/", pubHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
