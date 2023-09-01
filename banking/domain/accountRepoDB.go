package domain

import (
	"strconv"

	"github.com/H-proshanto/go-banking-microservice/banking/errs"
	"github.com/H-proshanto/go-banking-microservice/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepoDB struct {
	db *sqlx.DB
}

func NewAccountRepoDB(dbClient *sqlx.DB) *AccountRepoDB {
	return &AccountRepoDB{
		db: dbClient,
	}
}

func (r *AccountRepoDB) Save(account *Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES ($1, $2, $3, $4, $5) RETURNING account_id"

	var lastInsertID int
	err := r.db.Get(&lastInsertID, sqlInsert, account.CustomerID, account.OpeningDate, account.AccountType, account.Amount, account.Status)

	if err != nil {
		logger.Error("Error while creating an account" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	account.AccountID = strconv.Itoa(lastInsertID)

	return account, nil

}

func (r *AccountRepoDB) SaveTransaction(t *Transaction) (*Transaction, *errs.AppError) {
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("error while starting transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	var transactionID int64

	err = tx.QueryRow(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date)
			VALUES ($1, $2, $3, $4) RETURNING transaction_id`, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate).Scan(&transactionID)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}

	if t.IsWithDrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - $1 WHERE account_id = $2`, t.Amount, t.AccountID)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + $1 WHERE account_id = $2`, t.Amount, t.AccountID)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while updating account balance" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while committing transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}

	t.TransactionID = strconv.FormatInt(transactionID, 10)

	return t, nil

}

func (r *AccountRepoDB) FindBy(id *string) (*Account, *errs.AppError) {
	findCustomerByIdSqlQuery := "SELECT  customer_id, opening_date, account_type, amount, status FROM accounts WHERE account_id = $1"
	var c Account
	err := r.db.Get(&c, findCustomerByIdSqlQuery, id)

	if err != nil {
		logger.Error("Error while scanning customer" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return &c, nil
}
