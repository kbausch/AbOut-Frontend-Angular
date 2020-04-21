package endpoints

import "net/http"

// This is a helper function that sets the 'content type' and the 'access
// control allow origin' headers.
func corsAndContentHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

// OK sets the standard headers and writes status OK to the header.
func OK(w http.ResponseWriter) {
	corsAndContentHeaders(w)
	w.WriteHeader(http.StatusOK)
}

// CREATED sets the standard headers and writes status CREATED to the header.
func CREATED(w http.ResponseWriter) {
	corsAndContentHeaders(w)
	w.WriteHeader(http.StatusCreated)
}

// NOCONTENT sets the standard headers and writes status PUT to the header.
func NOCONTENT(w http.ResponseWriter) {
	corsAndContentHeaders(w)
	w.WriteHeader(http.StatusNoContent)
}
