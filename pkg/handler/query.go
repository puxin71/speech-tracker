package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

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

// Retrieve all the talks dataset and return it in JSON
func (q Query) GetAllTalks(w http.ResponseWriter, r *http.Request) {
	data, err := q.db.GetAllTalks()

	if err != nil {
		log.Println("fail to get all talks, error: ", err)
		updateHeader(w, err)
		return
	}

	writeJsonBody(w, data)
}

func (q Query) GetTalkAttendees(w http.ResponseWriter, r *http.Request) {
	var tkid int
	var err error

	// Extract the talk ID from the request
	vars := mux.Vars(r)
	if tkid, err = strconv.Atoi(strings.TrimSpace(vars["id"])); err != nil {
		log.Println("invalid talk id:", vars["id"], "err:", err)
		updateHeader(w, database.ErrInvalidTalkID)
		return
	}

	// Retrieve all the attendees that have registerted to the talk
	data, err := q.db.GetAttendees(tkid)
	if err != nil {
		log.Println("fail to query talk with id:", tkid)
		updateHeader(w, err)
		return
	}

	writeJsonBody(w, data)
}

// Populate response body with JSON formatted data
func writeJsonBody(w http.ResponseWriter, data interface{}) {
	payload, err := json.Marshal(data)
	if err != nil {
		log.Println("fail to encode data to the JSON format, error:", err)
		updateHeader(w, err)
		return
	}

	updateHeader(w, err)
	w.Write(payload)
}
