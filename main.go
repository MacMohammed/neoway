package main

import (
	"fmt"
	"log"
	"neoway/config"
	"neoway/upload"
	"net/http"
)

func main() {

	config.Carregar()

	mux := http.NewServeMux()
	mux.HandleFunc("/", upload.IndexHandler)
	mux.HandleFunc("/upload", upload.UploadFile)
	
	if err := http.ListenAndServe(fmt.Sprintf(":%d",config.Porta), mux); err != nil {
		log.Fatal(err)
	}
}
