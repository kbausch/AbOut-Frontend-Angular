package middleware

import (
	"net/http"
)

// Middleware defines the function signature for a middleware.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain applies middlewares to an http.Handler.
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
