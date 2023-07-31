package dto

import (
	"github.com/tricoman/banking/errs"
	"strings"
)

type NewTransactionRequest struct {
	CustomerId      string  `json:"customer_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (r NewTransactionRequest) Validate() *errs.AppError {
	if r.Amount < 0 {
		return errs.NewValidationError("Transaction amount can't be negative")
	}
	if strings.ToLower(r.TransactionType) != "deposit" && strings.ToLower(r.TransactionType) != "withdrawal" {
		return errs.NewValidationError("Transaction type should be 'deposit' or 'withdrawal'")
	}
	return nil
}
