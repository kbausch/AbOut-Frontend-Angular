package endpoints

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"gitlab.cs.mtech.edu/ESOF326/S20/AbOut/backend/middleware"
	"gitlab.cs.mtech.edu/ESOF326/S20/AbOut/backend/models"
)

// Type that implements the ProgramsRepository interface for mock testing.
var programsMockAPI = ProgramsAPI{mockRepo{}}

// GetPrograms writes a list of programs to w.
func (mockRepo) GetPrograms() (models.Programs, error) {
	programs := models.Programs{
		{
			Abbrev:          "CS",
			Name:            "Computer Science",
			CurrentSemester: "Spring 2020",
		},
		{
			Abbrev:          "SE",
			Name:            "Software Engineering",
			CurrentSemester: "Spring 2020",
		},
	}
	return programs, nil
}

func (mockRepo) DisassociateOutcome(a string, p string, i string) error {
	if a != "CS" || p != "CAC" || i != "1" {
		return fmt.Errorf("outcome association with prefix and identifier Doesnt exist: "+" %s and %s\n", p, i)
	}
	return nil
}

func TestGetPrograms_GoodCase_ResponseOK(t *testing.T) {
	// Arrange:
	// Create a request to send to the endpoint.
	req, err := http.NewRequest("GET", "/programs", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := http.HandlerFunc(programsMockAPI.GetPrograms)

	// Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check that the response body is as expected.
	expected := `[{"abbrev":"CS","name":"Computer Science","current_semester":"Spring 2020"},{"abbrev":"SE","name":"Software Engineering","current_semester":"Spring 2020"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	// Check that the headers are as expected.
	if ctype := rr.Header().Get(contentType); ctype != appjson {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, appjson)
	}
}

func TestDisassociateOutcome_GoodCase_ResponseOK(t *testing.T) {
	// Arrange:
	os.Setenv("authType", "jwt")
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test DisassociateOutcome.
	req, err := http.NewRequest("DELETE", "/programs/CS/outcomes/CAC/1", strings.NewReader("test body"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+fakeToken)
	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := mux.NewRouter()

	handler.HandleFunc("/programs/{program-abbrev}/outcomes/{prefix}/{identifier}", middleware.Chain(programsMockAPI.DisassociateOutcome, middleware.PermissionsMiddleware())).Methods("DELETE") // Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDisassociateOutcome_BadCase_AssociationDoesNotExist(t *testing.T) {
	// Arrange:
	os.Setenv("authType", "jwt")
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test DisassociateOutcome.
	req, err := http.NewRequest("DELETE", "/programs/CS/outcomes/CAC/2", strings.NewReader("test body"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+fakeToken)
	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := mux.NewRouter()

	handler.HandleFunc("/programs/{program-abbrev}/outcomes/{prefix}/{identifier}", middleware.Chain(programsMockAPI.DisassociateOutcome, middleware.PermissionsMiddleware())).Methods("DELETE") // Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestDisassociateOutcome_BadCase_NotAuthorized(t *testing.T) {
	// Arrange:
	os.Setenv("authType", "jwt")
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test DisassociateOutcome.
	req, err := http.NewRequest("DELETE", "/programs/CS/outcomes/CAC/1", strings.NewReader("test body"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+fakeTokenNonSuperuser)
	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := mux.NewRouter()

	handler.HandleFunc("/programs/{program-abbrev}/outcomes/{prefix}/{identifier}", middleware.Chain(programsMockAPI.DisassociateOutcome, middleware.PermissionsMiddleware())).Methods("DELETE")
	// Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusForbidden)
	}
}
