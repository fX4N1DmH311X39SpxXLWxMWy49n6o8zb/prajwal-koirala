package main

import (
	"log"
	"net/http"
	"os"
)

var err error

func main() {
	if folderExists("www") {
		http.Handle("/", http.FileServer(http.Dir("www")))
		err = http.ListenAndServe(":8080", nil)
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
