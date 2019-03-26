package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	_ = json.NewEncoder(res).Encode(req.Body)
}
