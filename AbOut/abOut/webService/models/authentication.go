package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// PermissionsTable represents the result of the fetch user stored procedure.
type PermissionsTable struct {
	IsSuperUser bool             `json:"is_super_user"`
	Rows        []PermissionsRow `json:"permissions"`
}

// PermissionsRow represents a row from the permissions table.
type PermissionsRow struct {
	ProgramName string `json:"program_name"`
	IsMananger  bool   `json:"is_manager"`
	IsObserver  bool   `json:"is_observer"`
}

// PermissionsRepository is the interface for permission handling inside the
// endpoints.
type PermissionsRepository interface {
	ReadHeaderData(*http.Request) error
	IsAuthenticated() bool
	IsSuperUser() bool
	GetProgramPermissions(string) (PermissionsRow, error)
}

// PermissionsHandler implements PermissionsRepository.
type PermissionsHandler struct {
	authenticated   bool
	permissionTable PermissionsTable
}

// ReadHeaderData reads permission JSON from the header into the
// PermissionsHandler struct for reading.
func (ph *PermissionsHandler) ReadHeaderData(r *http.Request) error {
	// Get the permissions string from the header.
	JSONPermissionsString := r.Header.Get("permissions")
	// Unmarshall the permissions from JSON into a PermissionTable.
	err := json.Unmarshal([]byte(JSONPermissionsString), &ph.permissionTable)
	if err != nil {
		return errors.New("failed to unmarshal permission header")
	}
	ph.authenticated = true
	return nil
}

// IsAuthenticated returns true if the user is authenticated.
func (ph PermissionsHandler) IsAuthenticated() bool {
	return ph.authenticated
}

// IsSuperUser returns if the user is a super user.
func (ph PermissionsHandler) IsSuperUser() bool {
	return ph.permissionTable.IsSuperUser
}

// GetProgramPermissions gets the permissions for a user from a program with
// the given name.
func (ph PermissionsHandler) GetProgramPermissions(programName string) (*PermissionsRow, error) {
	// n will be small but we can switch to a hashmap if we need to.
	for _, row := range ph.permissionTable.Rows {
		if row.ProgramName == programName {
			return &row, nil
		}
	}
	return nil, errors.New("permissions not found for this user given the program")
}

// PermissionsRepo implements PermissionsRepository for dependency injection.
type PermissionsRepo struct {
}

var repo PermissionsRepo

// GetPermissions fetches permissions from the database for the CAS username
// provided and returns it as raw JSON data.
func (PermissionsRepo) GetPermissions(usernameCAS string) (*PermissionsTable, error) {
	trx, _ := db.Begin()

	// Call the stored procedure to fetch all permission rows. Is super user result is stored in the is_super_user.
	rows, err := trx.Query("CALL permissions__fetch_user__sp(?, @is_super_user);", usernameCAS)
	if err != nil {
		return nil, err
	}
	permissionRows := []PermissionsRow{}
	for rows.Next() {
		var pRow PermissionsRow
		err = rows.Scan(&pRow.ProgramName, &pRow.IsMananger, &pRow.IsObserver)
		if err != nil {
			fmt.Println(err)
			continue
		}
		permissionRows = append(permissionRows, pRow)
	}
	rows.Close()

	// Select the is_super_user variable outputed by the stored procedure.
	var isSuperUser int
	superUserResult := trx.QueryRow("SELECT @is_super_user;")
	superUserResult.Scan(&isSuperUser)

	trx.Commit()
	return &PermissionsTable{IsSuperUser: isSuperUser != 0, Rows: permissionRows}, nil
}

// ParseJwtToken extracts the JWT from the authorization header and parses it into a jwt.Token object.
func ParseJwtToken(r *http.Request) (*jwt.Token, error) {
	// If there are no values associated with the key, Get returns "".
	// authorizationHeader should be of the form "Authorization: Bearer <token>".
	authorizationHeader := strings.Split(r.Header.Get("Authorization"), " ")
	if len(authorizationHeader) != 2 {
		return nil, errors.New("Authorization header of improper length")
	}

	tokenString := authorizationHeader[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// The token is invalid or not signed with RSA as expected.
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used to decode the token.
		return []byte(os.Getenv("secretKey")), nil
	})

	// Check if we failed to parse the token.
	if err != nil {
		return nil, err
	}

	return token, nil
}

// FetchPermissionsFromJwt takes a JWT token and returns a JSON byte slice of
// the user's permission data or an error.
func FetchPermissionsFromJwt(token *jwt.Token) ([]byte, error) {
	// Attempt to extract the claims from the token.
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// At this point, we have valid claims. We should fetch the user's
		// permissions information from the database and append it to the
		// request's headers.
		usernameCAS := fmt.Sprintf("%v", claims["usernameCAS"])

		// Fetch the user's permissions from the database.
		return FetchPermissionsFromCasUsername(usernameCAS)
	}
	return nil, errors.New("JWT claims not valid")
}

// FetchPermissionsFromCasUsername retrieves the user's permission in JSON from
// making a database query based upon their user name.
func FetchPermissionsFromCasUsername(usernameCAS string) ([]byte, error) {
	pTable, err := repo.GetPermissions(usernameCAS)
	if err != nil {
		// We failed to fetch permissions.
		// A user with no permissions may trigger this, so no logging.
		return nil, errors.New("failed to fetch permission data from the database")
	}

	// Convert the user's permissions into JSON.
	jsonData, err := json.Marshal(pTable)
	if err != nil {
		return nil, errors.New("Failed to marshall permission data")
	}
	return jsonData, nil
}
