package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var PathDirectory = `./StoreFile/`

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/documents", getDocuments).Methods("GET")
	router.HandleFunc("/documents/{ID}", getDocumentsById).Methods("GET")
	router.HandleFunc("/CreateDocuments", createDocument).Methods("POST")
	router.HandleFunc("/DeleteDocuments/{ID}", deleteDocument).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", router))
}
