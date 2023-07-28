package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/tricoman/banking/errs"
	"github.com/tricoman/banking/logger"
	"strconv"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	insertSql := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"
	result, err := d.client.Exec(insertSql, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
