package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getDocuments(w http.ResponseWriter, r *http.Request) {
	var docs []Document
	path := "./StoreFile/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, fil := range files {
		b, err := ioutil.ReadFile(path + fil.Name())
		if err != nil {
			fmt.Print(err)
		}
		fileContent := string(b)
		docs = append(docs, Document{ID: getMD5Checksum(fileContent), Name: fil.Name(), Size: fil.Size()})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)
}

func getMD5Checksum(content string) string {
	hasher := md5.New()
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil))
}
