package domain

import "github.com/H-proshanto/go-banking-microservice/banking/errs"

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}


type AccountRepo interface {
	Save(account *Account) (*Account, *errs.AppError) 
}
