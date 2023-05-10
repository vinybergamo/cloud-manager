package api

import (
	"encoding/json"
	"net/http"

	applications "github.com/vinybergamo/cloud-manager/apps"
)

var app App

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := applications.Create(app.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(res))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(res))
}

func apps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method == "DELETE" {
		res, err := applications.Destroy(app.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(res))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func deploy(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := applications.Deploy(app.Name, app.Git, app.Port)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(res))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(res))
}
