package api

import (
	"log"
	"net/http"

	"github.com/vinybergamo/cloud-manager/utils"
)

type App struct {
	Name string `json:"name"`
	Git  string `json:"git,omitempty"`
	Port string `json:"port,omitempty"`
}

func Init() {
	http.HandleFunc("/apps/create", create)
	http.HandleFunc("/apps", apps)
	http.HandleFunc("/apps/deploy", deploy)

	utils.Logger("green", "Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
