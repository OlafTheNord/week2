package main

import (
	"fmt"
	"log"
	"net/http"
	"week2/server"
)

func main() {
	routes := server.SetupRoutes()
	fmt.Println("server started on port 8000")

	log.Fatal(http.ListenAndServe(":8080", routes))
}
