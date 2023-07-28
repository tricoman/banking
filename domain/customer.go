package domain

import (
	"github.com/tricoman/banking/dto"
	"github.com/tricoman/banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c *Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c *Customer) AsDto() dto.CustomerResponse {
	statusAsText := c.statusAsText()
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      statusAsText,
	}
}

type CustomerRepository interface {
	FindBy(id string) (*Customer, *errs.AppError)
	FindAll() ([]Customer, *errs.AppError)
	FindAllActive() ([]Customer, *errs.AppError)
	FindAllInactive() ([]Customer, *errs.AppError)
}
