package handler

import (
	"fmt"
	"net/http"
)

type Home struct{}

func (Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home!")
}
