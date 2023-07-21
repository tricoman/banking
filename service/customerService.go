package service

import "github.com/tricoman/banking/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (service DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return service.repository.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repository: repository,
	}
}
