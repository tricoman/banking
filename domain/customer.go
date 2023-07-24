package domain

import "github.com/tricoman/banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindBy(id string) (*Customer, *errs.AppError)
}
