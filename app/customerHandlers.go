package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tricoman/banking/service"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(writer http.ResponseWriter, request *http.Request) {
	customers, err := ch.service.GetAllCustomers()

	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, customer)
	}

}

func writeResponse(writer http.ResponseWriter, code int, data interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		panic(err)
	}
}
