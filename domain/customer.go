package domain

import "github.com/tricoman/banking/errs"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindBy(id string) (*Customer, *errs.AppError)
	FindAll() ([]Customer, *errs.AppError)
	FindAllActive() ([]Customer, *errs.AppError)
	FindAllInactive() ([]Customer, *errs.AppError)
}
