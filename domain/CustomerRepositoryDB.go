package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tricoman/banking/errs"
	"log"
	"time"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (r CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	rows, err := r.client.Query(findAllSql)

	if err != nil {
		errorMessage := "Error while querying customer table :" + err.Error()
		log.Println(errorMessage)
		return nil, errs.NewUnexpectedError(errorMessage)
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var customer Customer
		err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.City,
			&customer.ZipCode,
			&customer.DateOfBirth,
			&customer.Status,
		)
		if err != nil {
			errorMessage := "Error while scanning customer rows :" + err.Error()
			log.Println(errorMessage)
			return nil, errs.NewUnexpectedError(errorMessage)
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (r CustomerRepositoryDB) FindBy(id string) (*Customer, *errs.AppError) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"
	row := r.client.QueryRow(findAllSql, id)

	var customer Customer
	err := row.Scan(
		&customer.Id,
		&customer.Name,
		&customer.City,
		&customer.ZipCode,
		&customer.DateOfBirth,
		&customer.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected DB error")
		}
	}
	return &customer, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	mysqlClient, err := sql.Open("mysql", "root:qwerty12345@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	mysqlClient.SetConnMaxLifetime(time.Minute * 3)
	mysqlClient.SetMaxOpenConns(10)
	mysqlClient.SetMaxIdleConns(10)
	return CustomerRepositoryDB{mysqlClient}
}
