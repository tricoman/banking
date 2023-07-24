package app

import (
	"github.com/gorilla/mux"
	"github.com/tricoman/banking/domain"
	"github.com/tricoman/banking/service"
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
	customerHandlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	router.HandleFunc(customersEndpoint, customerHandlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc(getCustomerEndpoint, customerHandlers.getCustomer).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
