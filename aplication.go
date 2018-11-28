package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func getDocuments(w http.ResponseWriter, r *http.Request) {
	var docs []Document
	docs = getDocumentsTypeList()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(docs)
}

func getDocumentsById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["ID"]
	doc, err := getDocumentByIdMd5(id)
	if err == nil {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(doc)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}
}

func getDocumentsTypeList() []Document {
	var docs []Document

	var fileList = getListFiles(PathDirectory)
	for _, value := range fileList {

		absolutePath := PathDirectory + value
		verifyValidPath(absolutePath)
		hashMd5, err := hashFileMd5(absolutePath)
		fileSize := getSizeFile(absolutePath)

		if err == nil {
			docs = append(docs,
				Document{ID: hashMd5, Name: value, Size: fileSize})
		}
	}

	return docs
}

func getDocumentByIdMd5(id string) (Document, error) {
	var docs []Document
	docs = getDocumentsTypeList()
	for _, doc := range docs {
		if doc.ID == id {
			return doc, nil
		}
	}
	return Document{}, errors.New("ID not found")
}

func createDocument(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	defer file.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errors.New("save to File in StoreFile")
	}
	f, err := os.OpenFile("./StoreFile/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errors.New("it was not saved")
	}
	defer f.Close()
	io.Copy(f, file)
	w.WriteHeader(http.StatusForbidden)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Successfully saved " + handler.Filename)
}

func deleteDocument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["ID"]
	var docs []Document
	docs = getDocumentsTypeList()
	sw := false
	docName := ""
	for _, doc := range docs {
		if doc.ID == id {
			sw = true
			docName = doc.Name
			os.Remove("./StoreFile/" + doc.Name)
		}
	}
	if !sw {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("File not found")
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("delete the file successfully " + docName)
	}

}

