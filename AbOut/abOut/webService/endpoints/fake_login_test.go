package endpoints

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var fakeLoginMockAPI = FakeLoginAPI{}

func TestGetToken_GoodCase_ResponseOK(t *testing.T) {
	// Arrange:
	// Create a request to send to the endpoint.
	req, err := http.NewRequest("GET", "/auth", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := http.HandlerFunc(fakeLoginMockAPI.GetToken)

	// Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// This test is interesting because we cannot test the validity of the token in
	// the response body as it is time dependent. The best we can do is see that the
	// token format is present.
	expected := `{"token"`
	split := strings.Split(rr.Body.String(), ":")[0]
	if split != expected {
		t.Errorf("handler returned unexpected body format: got %v want %v",
			split, expected)
	}
	// Check that the headers are as expected.
	if ctype := rr.Header().Get(contentType); ctype != appjson {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, appjson)
	}
}
