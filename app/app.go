package app

import (
	"log"
	"net/http"
)

const greetEndpoint = "/greet"
const getAllCustomersEndpoint = "/customers"

func Start() {
	defineRoutes()
	log.Fatal(startServer())
}

func defineRoutes() {
	http.HandleFunc(greetEndpoint, greetHandler)
	http.HandleFunc(getAllCustomersEndpoint, getAllCustomersHandler)
}

func startServer() error {
	return http.ListenAndServe("localhost:8000", nil)
}
