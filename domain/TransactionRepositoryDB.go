package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/tricoman/banking/errs"
	"github.com/tricoman/banking/logger"
	"strconv"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func (r TransactionRepositoryDB) Create(t Transaction) (*Transaction, float64, *errs.AppError) {
	newTransactionSql := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?,?,?,?)"
	result, err := r.client.Exec(newTransactionSql, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error while creating new transaction: " + err.Error())
		return nil, 0, errs.NewUnexpectedError("Unexpected error from DB")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new transaction: " + err.Error())
		return nil, 0, errs.NewUnexpectedError("Unexpected error from DB")
	}
	t.TransactionId = strconv.FormatInt(id, 10)
	finalBalance, appError := updateAccountBalance(t, r)
	if appError != nil {
		return nil, 0, appError
	}
	return &t, finalBalance, nil
}

func updateAccountBalance(t Transaction, r TransactionRepositoryDB) (float64, *errs.AppError) {
	accountBalance, appError := fetchAccountBalance(t, r)
	if appError != nil {
		return 0, appError
	}
	finalBalance := computingFinalBalance(t, accountBalance)
	if finalBalance < 0.0 {
		logger.Error("Error while withdrawing amount, insufficient account amount")
		return 0, errs.NewBadRequestError("Withdraw amount should be less than total account balance")
	}
	updateAccountSql := "UPDATE accounts SET amount = ? WHERE account_id = ?"
	_, err := r.client.Exec(updateAccountSql, finalBalance, t.AccountId)
	if err != nil {
		logger.Error("Error while updating account balance: " + err.Error())
		return 0, errs.NewUnexpectedError("Unexpected error from DB")
	}
	return finalBalance, nil
}

func computingFinalBalance(t Transaction, accountBalance float64) float64 {
	var finalBalance float64
	if t.TransactionType == "withdrawal" {
		finalBalance = accountBalance - t.Amount
	} else {
		finalBalance = accountBalance + t.Amount
	}
	return finalBalance
}

func fetchAccountBalance(t Transaction, r TransactionRepositoryDB) (float64, *errs.AppError) {
	selectAmountSql := "SELECT amount FROM accounts WHERE account_id = ?"
	var finalBalance []float64
	err := r.client.Select(&finalBalance, selectAmountSql, t.AccountId)
	if err != nil {
		logger.Error("Error while fetching account balance: " + err.Error())
		return 0, errs.NewUnexpectedError("Unexpected error from DB")
	}
	return finalBalance[0], nil
}

func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{dbClient}
}
