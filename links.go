package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (a *App) getLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid link ID")
		return
	}

	l := link{}
	l.getLink(id)
	respondWithJSON(w, http.StatusOK, l)
}

func (a *App) createLink(w http.ResponseWriter, r *http.Request) {
	var l link
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&l); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if !strings.HasPrefix(l.Target, "http") {
		l.Target = "http://" + l.Target
	}
	generatedId := RandStringBytes(6)
	for CheckAddressIsExists(generatedId) {
		generatedId = RandStringBytes(6)
	}
	l.Address = generatedId
	l.createLink()
	respondWithJSON(w, http.StatusCreated, l)
}

func (a *App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	l := link{}
	l.getTargetById(id)

	http.Redirect(w, r, l.Target, http.StatusMovedPermanently)
}
