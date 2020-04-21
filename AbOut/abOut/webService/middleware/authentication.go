package middleware

import (
	"fmt"
	"net/http"
	"os"

	"gitlab.cs.mtech.edu/ESOF326/S20/AbOut/backend/models"
	"gopkg.in/Dolphindalt/cas.v4"
)

// PermissionsMiddleware is a wrapper for all permissions options.
func PermissionsMiddleware() Middleware {
	authenticationType := os.Getenv("authType")
	if authenticationType == "jwt" {
		return jwtPermissions()
	}
	return casPermissions()
}

// jwtPermissions checks if a user is authenticated and appends their permission
// information to the request if so.
func jwtPermissions() Middleware {
	// Create a new middleware.
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the new HandlerFunc.
		return func(w http.ResponseWriter, r *http.Request) {
			// Now we do middleware things here.

			// Clear the permissions header for security.
			r.Header.Del("permissions")

			token, err := models.ParseJwtToken(r)

			// Check if we failed to parse the token.
			if err != nil {
				fmt.Printf("Error parsing token: %v\n", err)
				f(w, r)
				return
			}

			// Attempt to extract the claims from the token.
			jsonData, err := models.FetchPermissionsFromJwt(token)

			// Check if the token claims were valid.
			if err != nil {
				f(w, r)
				return
			}

			// Attach the permissions data as a request header.
			// This header is not allowed in the initial request.
			r.Header.Add("permissions", string(jsonData))

			// Call the next middleware or handler in the chain.
			f(w, r)
		}
	}
}

// casPermissions checks if the user is CAS authenticated and will append
// the user's permissions to the body of the request if so.
func casPermissions() Middleware {
	// Create a new middleware.
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the new HandlerFunc.
		return func(w http.ResponseWriter, r *http.Request) {
			// Now we do middleware things here.

			// Clear the permissions header for security.
			r.Header.Del("permissions")

			// Check if we are CAS authenticated.
			if cas.IsAuthenticated(r) {
				// We are authenticated.
				usernameCAS := cas.Username(r)

				// Now we fetch the user's permissions.
				jsonData, err := models.FetchPermissionsFromCasUsername(usernameCAS)

				// Check if the service token claims were valid.
				if err != nil {
					f(w, r)
					return
				}

				// Attach the permissions data as a request header.
				// This header is not allowed in the initial request.
				r.Header.Add("permissions", string(jsonData))
			} else {
				cas.RedirectToLogin(w, r)
				return
			}

			// Call the next middleware or handler in the chain.
			f(w, r)
		}
	}
}
