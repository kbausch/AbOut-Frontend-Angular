package server

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"

	"net/http"

	"gitlab.cs.mtech.edu/ESOF326/S20/AbOut/backend/endpoints"
	"gitlab.cs.mtech.edu/ESOF326/S20/AbOut/backend/middleware"
	"gitlab.cs.mtech.edu/ESOF326/S20/AbOut/backend/models"
)

// NewRouter creates a mux router and defines the routes for the application.
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	if os.Getenv("authType") == "jwt" {
		fAPI := endpoints.FakeLoginAPI{}
		r.HandleFunc("/auth/{usernameCAS}", fAPI.GetToken).Methods("GET")
	} else if os.Getenv("authType") == "cas" {
		lAPI := endpoints.CASLoginAPI{}
		r.HandleFunc("/auth/login", lAPI.Login).Methods("GET")
		r.HandleFunc("/auth/logout", lAPI.Logout).Methods("GET")
	}
	pAPI := endpoints.ProgramsAPI{Repo: models.ProgramsRepo{}}
	r.HandleFunc("/programs", pAPI.GetPrograms).Methods("GET")
	r.HandleFunc("/programs/{program-abbrev}/outcomes/{prefix}/{identifier}", middleware.Chain(pAPI.DisassociateOutcome, middleware.PermissionsMiddleware())).Methods("DELETE")

	oAPI := endpoints.OutcomesAPI{Repo: models.OutcomesRepo{}}
	r.HandleFunc("/outcomes", oAPI.GetOutcomes).Methods("GET")
	r.HandleFunc("/programs/{program-abbrev}/outcomes", oAPI.GetOutcomes).Methods("GET")
	r.HandleFunc("/outcomes/{prefix}/{identifier}", oAPI.GetOutcome).Methods("GET")
	r.HandleFunc("/outcomes/{prefix}/{identifier}", middleware.Chain(oAPI.CreateOutcome, middleware.PermissionsMiddleware())).Methods("POST")
	r.HandleFunc("/outcomes/{prefix}/{identifier}", middleware.Chain(oAPI.UpdateOutcome, middleware.PermissionsMiddleware())).Methods("PUT")
	r.HandleFunc("/outcomes/{prefix}/{identifier}", middleware.Chain(oAPI.DeleteOutcome, middleware.PermissionsMiddleware())).Methods("DELETE")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "reached home page") })
	return r
}
