package service

import (
	"github.com/tricoman/banking/domain"
	"github.com/tricoman/banking/errs"
)

type CustomerService interface {
	GetCustomer(id string) (*domain.Customer, *errs.AppError)
	GetAllCustomers() ([]domain.Customer, *errs.AppError)
	GetAllCustomersBy(status string) ([]domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repository.FindBy(id)
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	return s.repository.FindAll()
}

func (s DefaultCustomerService) GetAllCustomersBy(status string) ([]domain.Customer, *errs.AppError) {
	switch status {
	case "active":
		return s.repository.FindAllActive()
	case "inactive":
		return s.repository.FindAllInactive()
	default:
		return nil, errs.NewBadRequestError("Invalid status received: " + status)
	}
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repository: repository,
	}
}
