package domain

import (
	"github.com/tricoman/banking/dto"
	"github.com/tricoman/banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) AsResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
