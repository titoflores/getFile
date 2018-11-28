package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func hashFileMd5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil

}

func getListFiles(path string) []string {
	var files []string
	location, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range location {
		files = append(files, f.Name())
	}
	return files
}

func getSizeFile(path string) int64 {
	file, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	sizeFile := file.Size()
	return sizeFile
}
func verifyValidPath(absolutePath string) {
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		fmt.Println(err)
	}
}
