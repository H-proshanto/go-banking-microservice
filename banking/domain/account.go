package domain

import (
	"github.com/H-proshanto/go-banking-microservice/banking/dto"
	"github.com/H-proshanto/go-banking-microservice/banking/errs"
)

type Account struct {
	AccountID   string  `db:"account_id"`
	CustomerID  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a *Account) ToNewAccountResDTO() *dto.NewAccountRes {
	return &dto.NewAccountRes{
		AccountID: a.AccountID,
	}
}

type AccountRepo interface {
	Save(account *Account) (*Account, *errs.AppError)
	SaveTransaction(transaction *Transaction) (*Transaction, *errs.AppError)
	FindBy(accountID *string) (*Account, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount < amount
}
