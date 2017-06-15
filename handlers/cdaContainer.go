package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	db "../db"
	m "../models"
	"github.com/gorilla/mux"
)

//CdaContainerIndex root of the api.
func CdaContainerIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	cdaContainerArray, err := db.GetAll()
	if err = json.NewEncoder(w).Encode(cdaContainerArray); err != nil {
		panic(err)
	}
}

//CdaContainerSave create an CdaContainer.
func CdaContainerSave(w http.ResponseWriter, r *http.Request) {
	var cdaContainer m.CdaContainer
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &cdaContainer); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := db.Save(cdaContainer)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

//CdaContainerFindByID find by id.
func CdaContainerFindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cdaID := vars["cdaID"]
	var err error
	if len(cdaID) == 0 {
		panic(err)
	}
	cdaContainer, err := db.FindByID(cdaID)
	if len(cdaContainer.ID) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(cdaContainer); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(m.JSONErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

//CdaContainerUpdate find by id.
func CdaContainerUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cdaID := vars["cdaID"]
	var cdaContainer m.CdaContainer
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &cdaContainer); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := db.Update(cdaContainer, cdaID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	return
}

//CdaContainerDelete find by id.
func CdaContainerDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cdaID := vars["cdaID"]

	if len(cdaID) == 0 {
		var err error
		panic(err)
	}

	err := db.Remove(cdaID)

	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(m.JSONErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}
