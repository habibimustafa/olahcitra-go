package main

import (
	"bytes"
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
	router.HandleFunc("/inverse", ToInverse).Methods("POST")
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

func TestToInverse(t *testing.T) {
	data := [][]int{{55, 55, 55}, {75, 75, 75}, {125, 125, 125}}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("POST", "/inverse", bytes.NewBuffer(jsonData))
	res := httptest.NewRecorder()
	router().ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code, "Expect OK Response")

	var result Message
	_ = json.Unmarshal(res.Body.Bytes(), &result)

	assert.True(t, result.Success)
	assert.Equal(t, "Converting to inverse success", result.Message)

	var resBody [][]float64
	for _, p := range result.Body.([]interface{}) {
		var pixel []float64
		for _, q := range p.([]interface{}) {
			pixel = append(pixel, q.(float64))
		}
		resBody = append(resBody, pixel)
	}

	expectedBody := [][]float64{{200, 200, 200}, {180, 180, 180}, {130, 130, 130}}
	assert.Equal(t, expectedBody, resBody)
}
