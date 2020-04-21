package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var casAPI = CASLoginAPI{}

func TestLogin_RedirectCase_ResponseRedirect(t *testing.T) {
	// Arrange:
	// Create a request to send to the endpoint.
	req, err := http.NewRequest("GET", "/auth/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := http.HandlerFunc(casAPI.Login)

	// Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestLogout_RedirectCase_ResponseRedirect(t *testing.T) {
	// Arrange:
	// Create a request to send to the endpoint.
	req, err := http.NewRequest("GET", "/auth/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := http.HandlerFunc(casAPI.Logout)

	// Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}
