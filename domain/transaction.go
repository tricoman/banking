package domain

import (
	"github.com/tricoman/banking/dto"
	"github.com/tricoman/banking/errs"
)

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

func (t Transaction) AsResponseDto(updatedBalance float64) dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:  t.TransactionId,
		UpdatedBalance: updatedBalance,
	}
}

type TransactionRepository interface {
	Create(Transaction) (*Transaction, float64, *errs.AppError)
}
