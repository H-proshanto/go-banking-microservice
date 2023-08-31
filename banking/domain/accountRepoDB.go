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
	 sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES ($1, $2, $3, $4, $5)"

	 result, err := r.db.Exec(sqlInsert, account.AccountID, account.CustomerID, account.OpeningDate, account.AccountType, account.Amount, account.Status)

	 if err != nil {
		logger.Error("Error while creating an account" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	 }

	 id, err := result.LastInsertId()

	 if err != nil {
		logger.Error("Error while getting last insert id for new account" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	 }

	 account.AccountID = strconv.FormatInt(id, 10)

	 return account, nil
	}
