package endpoints

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const timeTokenIsValid = time.Hour

// FakeLoginAPI is a struct that will allow for dependency injection.
type FakeLoginAPI struct {
}

// GetToken returns a valid JWT token with some parameters.
func (api FakeLoginAPI) GetToken(w http.ResponseWriter, r *http.Request) {
	// Ensure the user is not already logged in
	if r.Header.Get("Authorization") != "" {
		http.Error(w, "You are already logged in", http.StatusBadRequest)
		return
	}

	usernameCAS := mux.Vars(r)["usernameCAS"]

	// Create a new token object, specifying method and claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usernameCAS": usernameCAS,
		// exp stands for expiration. It is the time until the token is no longer valid.
		"exp": time.Now().Add(timeTokenIsValid).Unix(),
	})

	// Sign and get the complete encoded token using the secret key.
	tokenString, err := token.SignedString([]byte(os.Getenv("secretKey")))

	if err != nil {
		log.Printf("Error after token encoding: %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	OK(w)
	w.Write([]byte("{\"token\":\"" + tokenString + "\"}"))
}
