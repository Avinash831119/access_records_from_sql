package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/search", retrieve).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	initializeRouter()
}
