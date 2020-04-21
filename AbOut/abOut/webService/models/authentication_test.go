package models

import (
	"log"
	"net/http"
	"testing"
)

func TestGetPermissions_ValidUser(t *testing.T) {
	// Arrange:
	usernameCAS := "xdolence"
	var pRepo PermissionsRepo

	// Act:
	// Perform the database query.
	pTable, err := pRepo.GetPermissions(usernameCAS)
	if err != nil {
		t.Errorf("Unexpected error querying for user permissions: %v", pTable)
	}

	// Assert:
	// Check that the response code is as expected.
	if pTable.IsSuperUser != false {
		t.Errorf("xaavan should not be a super user")
	}
	if len(pTable.Rows) != 3 {
		t.Errorf("expected 3 program permissions, got: %v", len(pTable.Rows))
	}
}

func TestGetPermissions_ValidSuperUser(t *testing.T) {
	// Arrange:
	usernameCAS := "jvesco"
	var pRepo PermissionsRepo

	// Act:
	// Perform the database query.
	pTable, err := pRepo.GetPermissions(usernameCAS)
	if err != nil {
		t.Errorf("Unexpected error querying for user permissions: %v", err)
	}

	// Assert:
	// Check that the response code is as expected.
	if pTable.IsSuperUser != true {
		t.Errorf("vesco should be a super user")
	}
}

func TestGetPermissions_InvalidUser(t *testing.T) {
	// Arrange:
	usernameCAS := "something absurd"
	var pRepo PermissionsRepo

	// Act:
	// Perform the database query.
	pTable, err := pRepo.GetPermissions(usernameCAS)
	if err != nil {
		t.Errorf("Unexpected error querying for user permissions: %v", pTable)
	}

	// Assert:
	// Check that the response code is as expected.
	if pTable.IsSuperUser != false {
		t.Errorf("a user who is not in the database should not be super user")
	}
	if len(pTable.Rows) != 0 {
		t.Errorf("expected 0 table rows, got: %v", len(pTable.Rows))
	}
}

func TestPermissionHandler_ReadHeaderDataValidHeader(t *testing.T) {
	// Arrange:
	// Build the request header with the permission data in it.
	jsonData := `{"is_super_user":true,"permissions":[{"program_name":"Computer Science","is_manager":true,"is_observer":false}]}`
	// Create a dummy request for passing to the RawHeaderData function.
	req, err := http.NewRequest("GET", "/whatever", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Set the permission header for the request data.
	req.Header.Set("permissions", jsonData)

	// Act:
	// Execute the function.
	var pHandler PermissionsHandler
	err = pHandler.ReadHeaderData(req)
	if err != nil {
		t.Fatalf("error occured reading header: %v\n", err)
	}
	log.Println(pHandler.IsSuperUser())

	// Assert:
	// Check that the header data was read properly.
	if pHandler.IsSuperUser() != true {
		t.Fatalf("expected IsSuperUser to be true, got false")
	}
	// Check that the user is authenticed.
	if pHandler.IsAuthenticated() != true {
		t.Fatalf("expected the user to be autheticated")
	}
	// Check that the program row was read properly.
	fetchedRow, err := pHandler.GetProgramPermissions("Computer Science")
	if err != nil {
		t.Fatalf("encountered an unexpected error fetching a permission row: %v\n", err)
	}
	// Test for expected permissions in the row
	if fetchedRow.IsMananger != true {
		t.Fatalf("expected the permission row to have IsManager set true")
	}
	if fetchedRow.IsObserver != false {
		t.Fatalf("expected the permission row to have IsObserver set false")
	}
}

func TestPermissionHandler_ReadHeaderDataEmptyHeader(t *testing.T) {
	// Arrange:
	// Build the request header with the permission data in it.
	jsonData := ``
	// Create a dummy request for passing to the RawHeaderData function.
	req, err := http.NewRequest("GET", "/whatever", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Set the permission header for the request data.
	req.Header.Set("permissions", jsonData)

	// Act:
	// Execute the function.
	var pHandler PermissionsHandler
	err = pHandler.ReadHeaderData(req)

	// Assert:
	// The user should not be authenticated.
	if pHandler.IsAuthenticated() == true {
		t.Fatalf("expected user to not be authenticated")
	}
	// We should get an error here.
	if err == nil {
		t.Fatalf("expected an error got none\n")
	}
}

func TestFetchPermissionsFromCasUsername_NoKnownUser(t *testing.T) {
	// Arrange:
	fakeUsernameCAS := "ligma"

	// Act:
	// Fetch permissions for the invalid user.
	data, err := FetchPermissionsFromCasUsername(fakeUsernameCAS)
	dataString := string(data)

	// Assert:
	// There should be no error but permissions data will be empty.
	if err != nil {
		t.Fatal("expected no errors when an invalid user was provided")
	}
	expected := `{"is_super_user":false,"permissions":[]}`
	if dataString != expected {
		t.Fatalf("wanted empty permissions JSON, got %v wanted %v",
			dataString, expected)
	}
}

func TestFetchPermissionsFromCasUsername_ValidKnownUser(t *testing.T) {
	// Arrange:
	usernameCAS := "xdolence"

	// Act:
	// We fetch the permissions of the user.
	data, err := FetchPermissionsFromCasUsername(usernameCAS)
	dataString := string(data)

	// Assert:
	// We should now have the user's permissions without errors.
	if err != nil {
		t.Fatal("expected no errors when an invalid user was provided")
	}
	expected := `{"is_super_user":false,"permissions":[{"program_name":"Software Engineering","is_manager":false,"is_observer":true},{"program_name":"Computer Science","is_manager":false,"is_observer":true},{"program_name":"Electrical Engineering","is_manager":false,"is_observer":false}]}`
	if dataString != expected {
		t.Fatalf("wanted empty permissions JSON, got %v wanted %v",
			dataString, expected)
	}
}
