package service

import (
	"github.com/tricoman/banking/domain"
	"github.com/tricoman/banking/dto"
	"github.com/tricoman/banking/errs"
	"time"
)

type TransactionService interface {
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

func (s DefaultTransactionService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	transaction := domain.Transaction{
		TransactionId:   "",
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	producedTransaction, updatedBalance, appError := s.repo.Create(transaction)
	if appError != nil {
		return nil, appError
	}
	response := producedTransaction.AsResponseDto(updatedBalance)

	return &response, nil
}

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func NewTransactionService(repo domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo}
}
