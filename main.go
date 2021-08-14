package main

import (
	"log"
	"neoway/upload"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", upload.IndexHandler)
	mux.HandleFunc("/upload", upload.UploadFile)
	
	if err := http.ListenAndServe(":4500", mux); err != nil {
		log.Fatal(err)
	}
}
