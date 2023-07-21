package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const greetEndpoint = "/greet"
const customersEndpoint = "/customers"
const getCustomerEndpoint = "/customers/{customer_id:[0-9]+}"

func Start() {
	startServer()
}

func startServer() {
	router := mux.NewRouter()
	router.HandleFunc(greetEndpoint, greetHandler).Methods(http.MethodGet)
	router.HandleFunc(customersEndpoint, getAllCustomersHandler).Methods(http.MethodGet)
	router.HandleFunc(customersEndpoint, createCustomerHandler).Methods(http.MethodPost)
	router.HandleFunc(getCustomerEndpoint, getCustomerById).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
