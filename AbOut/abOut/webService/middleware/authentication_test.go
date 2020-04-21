package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"gitlab.cs.mtech.edu/ESOF326/S20/AbOut/backend/endpoints"
)

const testSecretKey = "bruh420yoloswag#aa#bab#aaaa#a"
const fakeToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZUNBUyI6ImRjYXJvbiJ9.TkXWmK9kL2I4M2hR7a_-dFGsZVKxYUPmg2pjrlsE274"

// This dummyEndpoint does not do anything but is used to test middleware.
func dummyEndpoint(w http.ResponseWriter, r *http.Request) {
	// The permissions header of the request is written to the response body for testing purposes.
	w.Write([]byte(r.Header.Get("permissions")))
	endpoints.OK(w)
}

func TestJWTPermissionsMiddleware_GoodPermissions(t *testing.T) {
	// Arrange:
	// Set the secret key.
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test the middleware.
	req, err := http.NewRequest("GET", "/dummy", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Add the token for the authorization test.
	req.Header.Add("Authorization", "Bearer "+fakeToken)

	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	router := mux.NewRouter()
	router.HandleFunc("/dummy", Chain(dummyEndpoint, jwtPermissions())).Methods("GET")

	// Act:
	// Perform the request
	router.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check that the response body is as expected.
	expected := `{"is_super_user":false,"permissions":[{"program_name":"Electrical Engineering","is_manager":false,"is_observer":false}]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestJWTPermissionsMiddleware_NoPermissions(t *testing.T) {
	// Arrange:
	// Create a request to test the middleware.
	req, err := http.NewRequest("GET", "/dummy", bytes.NewReader([]byte(`{"some_json":"bruh"}`)))
	if err != nil {
		t.Fatal(err)
	}

	// We do not add the token: the user is not authneticated.

	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	router := mux.NewRouter()
	router.HandleFunc("/dummy", Chain(dummyEndpoint, jwtPermissions())).Methods("GET")

	// Act:
	// Perform the request
	router.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check that the response body is as expected.
	// The header will be an empty string.
	expected := ``
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCasMiddleware_NoServiceToken(t *testing.T) {
	// Arrange:
	// Create a request to test the middleware.
	req, err := http.NewRequest("GET", "/dummy", bytes.NewReader([]byte(`{"some_json":"bruh"}`)))
	if err != nil {
		t.Fatal(err)
	}

	// We do not add the service token: the user is not authneticated.

	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	router := mux.NewRouter()
	router.HandleFunc("/dummy", Chain(dummyEndpoint, casPermissions())).Methods("GET")

	// Act:
	// Perform the request
	router.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	// 500 because we are not CAS authneticated.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPermissionsMiddleware_JwtActive(t *testing.T) {
	// Arrange:
	// Set environment variable for jwt authentication.
	os.Setenv("authType", "jwt")

	// Act:
	// Try to use the middleware.
	var result = reflect.ValueOf(PermissionsMiddleware()).Pointer()

	// Assert:
	// The function returned and executed should be the jwtPermissions.
	var actual = reflect.ValueOf(jwtPermissions()).Pointer()
	if result != actual {
		t.Errorf("middleware returned wrong middleware got %v expected %v",
			result, actual)
	}
}

func TestPermissionsMiddleware_CasActive(t *testing.T) {
	// Arrange:
	// Set environment variable for cas authentication.
	os.Setenv("authType", "cas")

	// Act:
	// Try to use the middleware.
	var result = reflect.ValueOf(PermissionsMiddleware()).Pointer()

	// Assert:
	// The function returned and executed should be the casPermissions.
	var actual = reflect.ValueOf(casPermissions()).Pointer()
	if result != actual {
		t.Errorf("middleware returned wrong middleware got %v expected %v",
			result, actual)
	}
}
