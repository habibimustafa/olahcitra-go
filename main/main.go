package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"olahcitra/citra"
)

type Message struct {
	Success bool
	Message string
	Body    interface{}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
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
		pixel := citra.Pixel{elm[0], elm[1], elm[2]}
		_ = append(pixels, pixel)
	}

	gambar := citra.Citra{Data: pixels}

	_ = json.NewEncoder(res).Encode(gambar)
}
