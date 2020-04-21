package middleware

import (
	"github.com/rs/cors"
)

// CreateCorsMiddleware generates a middleware for handling CORs preflight requests.
func CreateCorsMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Location"},
		// Enable debugging for testing, consider disabling in production
		Debug: true,
	})
}
