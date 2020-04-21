package endpoints

import (
	"net/http"

	"gopkg.in/Dolphindalt/cas.v4"
)

// CASLoginAPI for dependency injection.
type CASLoginAPI struct{}

// Login will redirect the user to the CAS login or return OK if the user is
// authenticated.
func (api CASLoginAPI) Login(w http.ResponseWriter, r *http.Request) {
	// Ensure the user is not already logged in.
	if !cas.IsAuthenticated(r) {
		cas.RedirectToLogin(w, r)
		return
	}

	NOCONTENT(w)
}

// Logout will redirect the user to the CAS logout or return NOCONTENT if the
// user is not authenticated.
func (api CASLoginAPI) Logout(w http.ResponseWriter, r *http.Request) {
	// Ensure the user is not already logged in.
	if cas.IsAuthenticated(r) {
		cas.RedirectToLogout(w, r)
		return
	}

	NOCONTENT(w)
}
