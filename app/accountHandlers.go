package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tricoman/banking/dto"
	"github.com/tricoman/banking/service"
	"net/http"
)

type AccountHandlers struct {
	service service.AccountService
}

func (h AccountHandlers) newAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}
	if err == nil {
		request.CustomerId = customerId
		createAccount(w, request, h.service)
	}
}

func createAccount(w http.ResponseWriter, r dto.NewAccountRequest, s service.AccountService) {
	accountResponse, err := s.NewAccount(r)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusCreated, accountResponse)
	}
}
