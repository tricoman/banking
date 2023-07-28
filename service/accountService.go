package service

import (
	"github.com/tricoman/banking/domain"
	"github.com/tricoman/banking/dto"
	"github.com/tricoman/banking/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	account := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(account)

	if err != nil {
		return nil, err
	}
	response := newAccount.AsResponseDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
