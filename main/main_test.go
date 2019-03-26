package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomePage).Methods("GET")
	return router
}

func TestHomePage(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	router().ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code, "Expect OK Response")

	var resBody Message
	_ = json.Unmarshal(res.Body.Bytes(), &resBody)

	assert.True(t, resBody.Success)
	assert.Equal(t, "Welcome to Olah Citra", resBody.Message)
	assert.Nil(t, resBody.Body)
}
