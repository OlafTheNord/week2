package server

import (
	"log"
	"net/http"
)

func RunServer() {
	mux := http.NewServeMux()
	log.Print("server start")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
