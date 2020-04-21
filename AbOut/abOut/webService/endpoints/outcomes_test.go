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

const testSecretKey = "testSecretCodeForSuperUser"
const fakeToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZUNBUyI6Imp2ZXNjbyJ9.Ktl-NSZWD4Nu8Lx6ixIqc63xyprTyHIeKnkA6tbzE9k"
const fakeTokenNonSuperuser string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZUNBUyI6Inhkb2xlbmNlIn0.HiszsPKngJj4If5clGrSaJVQyryEpW11LPW2pbQ0WUQ"

// Type that implements the OutcomesRepository interface for mock testing.
var outcomesMockAPI = OutcomesAPI{mockRepo{}}

const (
	appjson     = "application/json; charset=utf-8"
	textplain   = "text/plain; charset=utf-8"
	contentType = "Content-Type"
)

// GetOutcomes writes a list of outcomes to w.
func (mockRepo) GetOutcomes() (models.Outcomes, error) {
	outcomes := models.Outcomes{
		{
			Prefix:     "CAC",
			Identifier: "1",
			Text:       "Students can adequetly bake a pie",
		},
		{
			Prefix:     "CAC",
			Identifier: "3",
			Text:       "Students full understand that the cake is a lie",
		},
	}
	return outcomes, nil
}

// GetOutcomesInProgram writes a list of outcomes to w.
func (mockRepo) GetOutcomesInProgram(a string) (models.Outcomes, error) {
	if a != "CS" {
		return models.Outcomes{}, fmt.Errorf("Program with abreviation %s not found", a)
	}
	outcomes := models.Outcomes{
		{
			Prefix:     "CAC",
			Identifier: "1",
			Text:       "Students can adequetly bake a pie",
		},
		{
			Prefix:     "CAC",
			Identifier: "3",
			Text:       "Students full understand that the cake is a lie",
		},
	}
	return outcomes, nil
}

// GetOutcome writes an outcome to w.
func (mockRepo) GetOutcome(p string, i string) (models.Outcome, error) {
	if p != "CAC" || i != "1" {
		return models.Outcome{}, fmt.Errorf("outcome with prefix and identifier not found: expected"+
			" %s and %s got %s and %s\n", "CAC", "1", p, i)
	}
	outcome := models.Outcome{
		Prefix:     "CAC",
		Identifier: "1",
		Text:       "an ability to analyze a complex computing problem and to apply principles of computing and other relevant disciplines to identify solutions",
	}
	return outcome, nil
}

// CreateOutcome adds an outcome to the database.
func (mockRepo) CreateOutcome(p string, i string, c string) error {
	if p == "CAC" && i == "1" {
		return fmt.Errorf("outcome with prefix and identifier exists: "+" %s and %s\n", p, i)
	}
	return nil
}

// DeleteOutcome writes an outcome to w.
func (mockRepo) DeleteOutcome(p string, i string) error {
	if p != "CAC" || i != "1" {
		return fmt.Errorf("outcome with prefix and identifier Doesnt exist: "+" %s and %s\n", p, i)
	}
	return nil
}

// UpdateOutcome updates an outcome in the database.
func (mockRepo) UpdateOutcome(p string, i string, c string) error {
	if p != "CAC" || i != "1" {
		return fmt.Errorf("outcome with prefix and identifier doesnt exist: "+" %s and %s\n", p, i)
	}
	return nil
}

func TestGetOutcomes_GoodCase_ResponseOK(t *testing.T) {
	// Arrange:
	// Create a request to send to the endpoint.
	req, err := http.NewRequest("GET", "/outcomes", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := http.HandlerFunc(outcomesMockAPI.GetOutcomes)

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
	expected := `[{"prefix":"CAC","identifier":"1","text":"Students can adequetly bake a pie","begin":"","end":""},{"prefix":"CAC","identifier":"3","text":"Students full understand that the cake is a lie","begin":"","end":""}]`
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

func TestGetOutcomesInProgram_GoodCase_ResponseOK(t *testing.T) {
	// Arrange:
	// Create a request to send to the endpoint.
	req, err := http.NewRequest("GET", "/programs/CS/outcomes", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Create a router to route the request and url parameters.
	handler := mux.NewRouter()
	// Set the handler for the test.
	handler.HandleFunc("/programs/{program-abbrev}/outcomes", outcomesMockAPI.GetOutcomesInProgram)

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
	expected := `[{"prefix":"CAC","identifier":"1","text":"Students can adequetly bake a pie","begin":"","end":""},{"prefix":"CAC","identifier":"3","text":"Students full understand that the cake is a lie","begin":"","end":""}]`
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
func TestGetOutcome_GoodCase_ResponseOK(t *testing.T) {
	// Arrange:
	// Create a request to send to the endpoint.
	req, err := http.NewRequest("GET", "/outcome/CAC/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Create a router to route the request and url parameters.
	handler := mux.NewRouter()
	// Set the handler for the test.
	handler.HandleFunc("/outcome/{prefix}/{identifier}", outcomesMockAPI.GetOutcome)

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
	expected := `{"prefix":"CAC","identifier":"1","text":"an ability to analyze a complex computing problem and to apply principles of computing and other relevant disciplines to identify solutions","begin":"","end":""}`
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

func TestCreateOutcome_GoodCase_ResponseOK(t *testing.T) {
	// Arrange:
	os.Setenv("authType", "jwt")
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test CreateOutcome.
	req, err := http.NewRequest("POST", "/outcome/CAC/2", strings.NewReader("test body"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+fakeToken)
	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := mux.NewRouter()

	handler.HandleFunc("/outcome/{prefix}/{identifier}", middleware.Chain(outcomesMockAPI.CreateOutcome, middleware.PermissionsMiddleware())).Methods("POST")
	// Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateOutcome_BadCase_NotAuthorized(t *testing.T) {
	// Arrange:
	os.Setenv("authType", "jwt")
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test CreateOutcome.
	req, err := http.NewRequest("POST", "/outcome/CAC/2", strings.NewReader("test body"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+fakeTokenNonSuperuser)
	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := mux.NewRouter()

	handler.HandleFunc("/outcome/{prefix}/{identifier}", middleware.Chain(outcomesMockAPI.CreateOutcome, middleware.PermissionsMiddleware())).Methods("POST")
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

func TestCreateOutcome_BadCase_OutComeExists(t *testing.T) {
	// Arrange:
	os.Setenv("authType", "jwt")
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test CreateOutcome.
	req, err := http.NewRequest("POST", "/outcome/CAC/1", strings.NewReader("test body"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+fakeToken)
	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := mux.NewRouter()

	handler.HandleFunc("/outcome/{prefix}/{identifier}", middleware.Chain(outcomesMockAPI.CreateOutcome, middleware.PermissionsMiddleware())).Methods("POST")
	// Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusConflict)
	}
}

func TestDeleteOutcome_GoodCase_ResponseOK(t *testing.T) {
	// Arrange:
	os.Setenv("authType", "jwt")
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test DeleteOutcome.
	req, err := http.NewRequest("DELETE", "/outcome/CAC/1", strings.NewReader("test body"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+fakeToken)
	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := mux.NewRouter()

	handler.HandleFunc("/outcome/{prefix}/{identifier}", middleware.Chain(outcomesMockAPI.DeleteOutcome, middleware.PermissionsMiddleware())).Methods("DELETE")
	// Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteOutcome_BadCase_NotAuthorized(t *testing.T) {
	// Arrange:
	os.Setenv("authType", "jwt")
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test DeleteOutcome.
	req, err := http.NewRequest("DELETE", "/outcome/CAC/1", strings.NewReader("test body"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+fakeTokenNonSuperuser)
	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := mux.NewRouter()

	handler.HandleFunc("/outcome/{prefix}/{identifier}", middleware.Chain(outcomesMockAPI.DeleteOutcome, middleware.PermissionsMiddleware())).Methods("DELETE")
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

func TestDeleteOutcome_BadCase_OutComeHasDependencies(t *testing.T) {
	// Arrange:
	os.Setenv("authType", "jwt")
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test DeleteOutcome.
	req, err := http.NewRequest("DELETE", "/outcome/CAC/2", strings.NewReader("test body"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+fakeToken)
	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := mux.NewRouter()

	handler.HandleFunc("/outcome/{prefix}/{identifier}", middleware.Chain(outcomesMockAPI.DeleteOutcome, middleware.PermissionsMiddleware())).Methods("DELETE")
	// Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestUpdateOutcome_GoodCase_ResponseOK(t *testing.T) {
	// Arrange:
	os.Setenv("authType", "jwt")
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test UpdateOutcome.
	req, err := http.NewRequest("PUT", "/outcome/CAC/1", strings.NewReader("test body"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+fakeToken)
	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := mux.NewRouter()

	handler.HandleFunc("/outcome/{prefix}/{identifier}", middleware.Chain(outcomesMockAPI.UpdateOutcome, middleware.PermissionsMiddleware())).Methods("PUT")
	// Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateOutcome_BadCase_NotAuthorized(t *testing.T) {
	// Arrange:
	os.Setenv("authType", "jwt")
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test UpdateOutcome.
	req, err := http.NewRequest("PUT", "/outcome/CAC/1", strings.NewReader("test body"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+fakeTokenNonSuperuser)
	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := mux.NewRouter()

	handler.HandleFunc("/outcome/{prefix}/{identifier}", middleware.Chain(outcomesMockAPI.UpdateOutcome, middleware.PermissionsMiddleware())).Methods("PUT")
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

func TestUpdateOutcome_BadCase_OutComeDoesntExist(t *testing.T) {
	// Arrange:
	os.Setenv("authType", "jwt")
	os.Setenv("secretKey", testSecretKey)

	// Create a request to test UpdateOutcome.
	req, err := http.NewRequest("PUT", "/outcome/CAC/2", strings.NewReader("test body"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+fakeToken)
	// Create a response recorder to capture the response.
	rr := httptest.NewRecorder()
	// Set the handler for the test.
	handler := mux.NewRouter()

	handler.HandleFunc("/outcome/{prefix}/{identifier}", middleware.Chain(outcomesMockAPI.UpdateOutcome, middleware.PermissionsMiddleware())).Methods("PUT")
	// Act:
	// Perform the request.
	handler.ServeHTTP(rr, req)

	// Assert:
	// Check that the response code is as expected.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
