package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vinybergamo/cloud-manager/apps"
)

func createApp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newApp App
	err := json.NewDecoder(r.Body).Decode(&newApp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Println(newApp.Name)

	res, err := apps.Create(newApp.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(res))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(res))
}
