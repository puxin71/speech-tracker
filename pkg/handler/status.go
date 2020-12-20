package handler

import (
	"net/http"

	"github.com/puxin71/talk-server/pkg/database"
)

// Map the internal server error to the proper http.statusCode
func updateHeader(w http.ResponseWriter, err error) {
	var statusCode int

	switch err {
	case nil:
		statusCode = http.StatusOK
	case database.ErrInvalidTalkID:
		statusCode = http.StatusNotFound
	default:
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
}
