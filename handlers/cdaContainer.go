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
//CdaContainerFindByID find by id.
func CdaContainerFindByNHC(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nhc := vars["NHC"]
	var err error
	if len(nhc) == 0 {
		panic(err)
	}
	q := `{ "nhc:"` + nhc + `, "episodeType:{"$ne": "H"}"}`
	query, _ := json.Marshal(q)

	fmt.Println(string(query))

	cdaContainerArray, err := db.FindByQuery(1, query)
	if len(cdaContainer.ID) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(cdaContainerArray); err != nil {
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

}
