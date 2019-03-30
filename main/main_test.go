package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createRequest(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, body)
	res := httptest.NewRecorder()
	router().ServeHTTP(res, req)

	return res
}

func getResultBody(result Message) [][]float64 {
	var resBody [][]float64
	for _, p := range result.Body.([]interface{}) {
		var pixel []float64
		for _, q := range p.([]interface{}) {
			pixel = append(pixel, q.(float64))
		}
		resBody = append(resBody, pixel)
	}

	return resBody
}

func getResult(res *httptest.ResponseRecorder) Message {
	var resBody Message
	_ = json.Unmarshal(res.Body.Bytes(), &resBody)
	return resBody
}

func TestHomePage(t *testing.T) {
	res := createRequest("GET", "/", nil)
	assert.Equal(t, 200, res.Code, "Expect OK Response")
	result := getResult(res)

	assert.True(t, result.Success)
	assert.Equal(t, "Welcome to Olah Citra", result.Message)
	assert.Nil(t, result.Body)
}

func TestToInverse(t *testing.T) {
	data := [][]int{{55, 55, 55}, {75, 75, 75}, {125, 125, 125}}
	jsonData, _ := json.Marshal(&data)

	res := createRequest("POST", "/inverse", bytes.NewBuffer(jsonData))
	assert.Equal(t, 200, res.Code, "Expect OK Response")

	result := getResult(res)
	assert.True(t, result.Success)
	assert.Equal(t, "Converting to inverse success", result.Message)

	resBody := getResultBody(result)
	expectedBody := [][]float64{{200, 200, 200}, {180, 180, 180}, {130, 130, 130}}
	assert.Equal(t, expectedBody, resBody)
}

func TestToGray(t *testing.T) {
	data := [][]int{{0, 25, 50}, {75, 100, 125}, {150, 175, 200}}
	jsonData, _ := json.Marshal(&data)

	res := createRequest("POST", "/grayscale", bytes.NewBuffer(jsonData))
	assert.Equal(t, 200, res.Code, "Expect OK Response")

	result := getResult(res)
	assert.True(t, result.Success)
	assert.Equal(t, "Converting to grayscale success", result.Message)

	resBody := getResultBody(result)
	for _, pixel := range resBody {
		for j := range pixel {
			k := j + 1
			if k == len(pixel) {
				k = 0
			}

			assert.Equal(t, pixel[j], pixel[k])
		}
	}
}

func TestToBinary(t *testing.T) {
	data := [][]int{{0, 25, 50}, {75, 100, 125}, {150, 175, 200}, {225, 250, 255}}
	jsonData, _ := json.Marshal(&data)

	res := createRequest("POST", "/binary", bytes.NewBuffer(jsonData))
	assert.Equal(t, 200, res.Code, "Expect OK Response")

	result := getResult(res)
	assert.True(t, result.Success)
	assert.Equal(t, "Converting to binary success", result.Message)

	actual := getResultBody(result)
	assert.Equal(t, 4, len(actual))

	var expected [][]float64
	for _, px := range data {
		var pixel []float64
		for _, e := range px {
			ex := e
			if ex > 128 {
				ex = 255
			} else {
				ex = 0
			}
			pixel = append(pixel, float64(ex))
		}
		expected = append(expected, pixel)
	}

	assert.Equal(t, expected, actual)
}
