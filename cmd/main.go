package main

import (
	v1 "go-api/api/v1"
	"go-api/config"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	router := v1.NewRouter()

	log.Println("Starting Server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
