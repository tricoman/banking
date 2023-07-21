package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

func greetHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World!!")
}

func getAllCustomersHandler(writer http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{"Arnau", "Sant Feliu de codines", "08182"},
		{"Perico", "Barcelona", "08182"},
		{"Palotes", "Vic", "08182"},
	}

	if request.Header.Get("Content-Type") == "application/xml" {
		writer.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(writer).Encode(customers)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(customers)
	}
}

func createCustomerHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Customer creation request received")

}

func getCustomerById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprint(writer, vars["customer_id"])
}
