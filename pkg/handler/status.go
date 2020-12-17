package handler

import "net/http"

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
