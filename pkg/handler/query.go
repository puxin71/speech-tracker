package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/puxin71/talk-server/pkg/database"
)

// Query service which uses the DB service to retrive the dataset
type Query struct {
	db database.DB
}

// Instantiate a query instance
func NewQuery(db database.DB) Query {
	return Query{db: db}
}

func updateHeader(w http.ResponseWriter, err error) {
	var statusCode int

	switch err {
	case nil:
		statusCode = http.StatusOK
	default:
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
}

// Retrieve all the talks dataset and return it in JSON
func (q Query) GetAllTalks(w http.ResponseWriter, r *http.Request) {
	talks, err := q.db.GetAllTalks()

	if err != nil {
		log.Println("fail to get all talks, error: ", err)
		updateHeader(w, err)
		return
	}

	payload, err := json.Marshal(talks)
	if err != nil {
		log.Println("fail to encode data to the JSON format, error: ", err)
		updateHeader(w, err)
		return
	}

	updateHeader(w, err)
	w.Write(payload)
}

func (q Query) GetTalk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	updateHeader(w, nil)
	fmt.Fprintf(w, "Handler, GetTalk with vars: %s", vars["id"])
}
