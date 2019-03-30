package main

import (
	"log"
	"net/http"
)

type Message struct {
	Success bool
	Message string
	Body    interface{}
}

func main() {
	log.Fatal(http.ListenAndServe(":8000", router()))
}
