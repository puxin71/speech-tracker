package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func updateHeader(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

func GetAllTalks(w http.ResponseWriter, r *http.Request) {
	updateHeader(w, http.StatusOK)
	fmt.Fprintf(w, "Handler, GetAllTalks")
}

func GetTalk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	updateHeader(w, http.StatusOK)
	fmt.Fprintf(w, "Handler, GetTalk with vars: %s", vars["id"])
}
