package service

import (
	"github.com/tricoman/banking/domain"
	"github.com/tricoman/banking/dto"
	"github.com/tricoman/banking/errs"
)

type CustomerService interface {
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetAllCustomersBy(status string) ([]dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repository.FindBy(id)
	if err != nil {
		return nil, err
	}
	response := c.AsDto()
	return &response, nil
}

func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repository.FindAll()
	var responses []dto.CustomerResponse

	for i, c := range customers {
		responses[i] = c.AsDto()
	}

	return responses, err
}

func (s DefaultCustomerService) GetAllCustomersBy(status string) ([]dto.CustomerResponse, *errs.AppError) {
	var (
		domainCustomers []domain.Customer
		customerDtos    []dto.CustomerResponse
		err             *errs.AppError
	)
	switch status {
	case "active":
		domainCustomers, err = s.repository.FindAllActive()
	case "inactive":
		domainCustomers, err = s.repository.FindAllInactive()
	default:
		return nil, errs.NewBadRequestError("Invalid status received: " + status)
	}

	if err != nil {
		return nil, err
	}

	for i, c := range domainCustomers {
		customerDtos[i] = c.AsDto()
	}

	return customerDtos, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repository: repository,
	}
}
