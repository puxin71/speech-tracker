package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/puxin71/talk-server/pkg"

	"github.com/gorilla/mux"
	"github.com/puxin71/talk-server/pkg/handler"
	"github.com/stretchr/testify/assert"
)

var router *mux.Router

func init() {
	router = handler.NewRouter(pkg.MockConfigProvider{})
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func TestGetAllTalks(t *testing.T) {
	req, _ := http.NewRequest("GET", "/talks", nil)
	recorder := executeRequest(req)
	resp := recorder.Result()

	_, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-type"))
	//assert.Equal(t, "Handler, GetAllTalks", string(body))
}

func TestGetTalkAttendees(t *testing.T) {
	req, _ := http.NewRequest("GET", "/talks/0/attendees", nil)
	recorder := executeRequest(req)
	resp := recorder.Result()

	_, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-type"))
}
