package main

import (
	"log"
	"net/http"
	"os"
)

var (
	err        error
	filePath   = "www"
	serverPort = ":8080"
)

func main() {
	if folderExists(filePath) {
		http.Handle("/", http.FileServer(http.Dir(filePath)))
		err = http.ListenAndServe(serverPort, nil)
		handleError(err)
	} else {
		log.Println("Error: Failed to find www folder")
	}
}

func handleError(err error) {
	if err != nil {
		log.Print(err)
	}
}

func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
