package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/tricoman/banking/errs"
	"github.com/tricoman/banking/logger"
	"time"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (r CustomerRepositoryDB) FindBy(id string) (*Customer, *errs.AppError) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"
	var customer Customer
	err := r.client.Get(&customer, findAllSql, id)

	if err != nil && err == sql.ErrNoRows {
		return nil, errs.NewNotFoundError("customer not found")
	}

	if err != nil {
		logger.Error("Error scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected DB error")
	}

	return &customer, nil
}

func (r CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	return queryDBForCustomers(r, findAllSql)
}

func (r CustomerRepositoryDB) FindAllActive() ([]Customer, *errs.AppError) {
	findActiveSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = 1"
	return queryDBForCustomers(r, findActiveSQL)
}

func (r CustomerRepositoryDB) FindAllInactive() ([]Customer, *errs.AppError) {
	findInactiveSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = 0"
	return queryDBForCustomers(r, findInactiveSQL)
}

func queryDBForCustomers(r CustomerRepositoryDB, query string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	err := r.client.Select(&customers, query)
	if err != nil {
		errorMessage := "Error while querying customer table :" + err.Error()
		logger.Error(errorMessage)
		return nil, errs.NewUnexpectedError(errorMessage)
	}
	return customers, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	mysqlClient, err := sqlx.Open("mysql", "root:qwerty12345@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	mysqlClient.SetConnMaxLifetime(time.Minute * 3)
	mysqlClient.SetMaxOpenConns(10)
	mysqlClient.SetMaxIdleConns(10)
	return CustomerRepositoryDB{mysqlClient}
}
