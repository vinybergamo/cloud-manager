package api

import (
	"log"
	"net/http"
)

type App struct {
	Name string `json:"name"`
}

func Init() {
	http.HandleFunc("/apps/create", createApp)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
