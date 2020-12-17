package handler

import (
	"github.com/gorilla/mux"
	"github.com/puxin71/talk-server/pkg/middleware"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/talks", GetAllTalks).Methods("GET")
	router.HandleFunc("/talks/{id:[0-9]+}", GetTalk).Methods("GET")

	router.HandleFunc("/", Home)

	router.Use(middleware.Logger)
	return router
}
