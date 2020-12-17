package handler

import (
	"github.com/gorilla/mux"
	"github.com/puxin71/talk-server/pkg"
	"github.com/puxin71/talk-server/pkg/database"
	"github.com/puxin71/talk-server/pkg/middleware"
)

func NewRouter(config pkg.ConfigProvider) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/talks", NewQuery(database.NewFileLoader(config)).GetAllTalks).Methods("GET")
	router.HandleFunc("/talks/{id:[0-9]+}", NewQuery(database.NewFileLoader(config)).GetTalk).Methods("GET")

	router.HandleFunc("/", Home)

	router.Use(middleware.Logger)
	return router
}
