package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/documents", getDocuments).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
