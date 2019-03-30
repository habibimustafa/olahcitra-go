package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"olahcitra/citra"
)

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomePage).Methods("GET")
	router.HandleFunc("/inverse", ToInverse).Methods("POST")
	router.HandleFunc("/grayscale", ToGrayScale).Methods("POST")
	return router
}

func HomePage(res http.ResponseWriter, req *http.Request) {
	returnData := Message{true, "Welcome to Olah Citra", nil}
	_ = json.NewEncoder(res).Encode(returnData)
}

func ToInverse(res http.ResponseWriter, req *http.Request) {
	var jsonParam [][]float64

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&jsonParam); err != nil {
		returnData := Message{false, "Failed decoding data", nil}
		_ = json.NewEncoder(res).Encode(returnData)
		return
	}

	defer req.Body.Close()

	var pixels []citra.Pixel
	for _, elm := range jsonParam {
		pixel := citra.Pixel{Red: elm[0], Green: elm[1], Blue: elm[2]}
		pixels = append(pixels, pixel)
	}

	img := citra.Citra{Data: pixels}
	inv := img.Inverse()

	returnData := Message{true, "Converting to inverse success", inv.Array()}
	_ = json.NewEncoder(res).Encode(returnData)
}

func ToGrayScale(res http.ResponseWriter, req *http.Request) {
	var jsonParam [][]float64

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&jsonParam); err != nil {
		returnData := Message{false, "Failed decoding data", nil}
		_ = json.NewEncoder(res).Encode(returnData)
		return
	}

	defer req.Body.Close()

	var pixels []citra.Pixel
	for _, elm := range jsonParam {
		pixel := citra.Pixel{Red: elm[0], Green: elm[1], Blue: elm[2]}
		pixels = append(pixels, pixel)
	}

	img := citra.Citra{Data: pixels}
	img.Gray()

	returnData := Message{true, "Converting to inverse success", img.Array()}
	_ = json.NewEncoder(res).Encode(returnData)
}
