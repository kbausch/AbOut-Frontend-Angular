package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.cs.mtech.edu/ESOF326/S20/AbOut/backend/models"
)

// ProgramsAPI is a struct that will allow for dependency injection.
type ProgramsAPI struct {
	Repo models.ProgramsRepository
}

// GetPrograms writes a list of programs to w.
func (api ProgramsAPI) GetPrograms(w http.ResponseWriter, r *http.Request) {
	o, err := api.Repo.GetPrograms()
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(o)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	OK(w)
	w.Write(b)
}

// DisassociateOutcome removes a assosiation between a program and outcome
func (api ProgramsAPI) DisassociateOutcome(w http.ResponseWriter, r *http.Request) {
	var pHandler models.PermissionsHandler
	err := pHandler.ReadHeaderData(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	log.Println(pHandler)
	if pHandler.IsSuperUser() != true {
		http.Error(w, "", http.StatusForbidden)
		return
	}

	prefix := mux.Vars(r)["prefix"]
	identifier := mux.Vars(r)["identifier"]
	abrev := mux.Vars(r)["program-abbrev"]
	err = api.Repo.DisassociateOutcome(abrev, prefix, identifier)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}
}
