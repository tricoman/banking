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
		h.createAccount(w, request)
	}
}

func (h AccountHandlers) createAccount(w http.ResponseWriter, r dto.NewAccountRequest) {
	account, err := h.service.NewAccount(r)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusCreated, account)
	}
}
