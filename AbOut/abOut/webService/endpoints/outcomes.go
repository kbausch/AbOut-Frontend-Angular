package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.cs.mtech.edu/ESOF326/S20/AbOut/backend/models"
)

// OutcomesAPI is a struct that will allow for dependency injection.
type OutcomesAPI struct {
	Repo models.OutcomesRepository
}

// GetOutcomes writes a list of outcomes to w.
func (api OutcomesAPI) GetOutcomes(w http.ResponseWriter, r *http.Request) {
	o, err := api.Repo.GetOutcomes()
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

// GetOutcomesInProgram writes a list of outcomes to w.
func (api OutcomesAPI) GetOutcomesInProgram(w http.ResponseWriter, r *http.Request) {
	abbrev := mux.Vars(r)["program-abbrev"]
	o, err := api.Repo.GetOutcomesInProgram(abbrev)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	b, err := json.Marshal(o)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	OK(w)
	w.Write(b)

}

// GetOutcome writes a outcome to w.
func (api OutcomesAPI) GetOutcome(w http.ResponseWriter, r *http.Request) {
	prefix := mux.Vars(r)["prefix"]
	identifier := mux.Vars(r)["identifier"]

	o, err := api.Repo.GetOutcome(prefix, identifier)
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

//CreateOutcome adds a defined outcome to the database
func (api OutcomesAPI) CreateOutcome(w http.ResponseWriter, r *http.Request) {
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
	bodyB, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	body := string(bodyB)

	err = api.Repo.CreateOutcome(prefix, identifier, body)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusConflict)
		return
	}
}

//DeleteOutcome removes a defined outcome from the database
func (api OutcomesAPI) DeleteOutcome(w http.ResponseWriter, r *http.Request) {
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

	err = api.Repo.DeleteOutcome(prefix, identifier)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

//UpdateOutcome updates a defined outcome in the database
func (api OutcomesAPI) UpdateOutcome(w http.ResponseWriter, r *http.Request) {
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
	bodyB, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	body := string(bodyB)

	err = api.Repo.UpdateOutcome(prefix, identifier, body)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}
}
