package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tricoman/banking/dto"
	"github.com/tricoman/banking/service"
	"net/http"
)

type TransactionHandlers struct {
	service service.TransactionService
}

func (h TransactionHandlers) newTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]
	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}
	if err == nil {
		request.CustomerId = customerId
		request.AccountId = accountId
		createTransaction(w, request, h.service)
	}
}

func createTransaction(w http.ResponseWriter, r dto.NewTransactionRequest, s service.TransactionService) {
	transactionResponse, err := s.NewTransaction(r)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusCreated, transactionResponse)
	}
}
