package domain

import "github.com/H-proshanto/go-banking-microservice/banking/dto"

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionID   string  `db:"transaction_id"`
	AccountID       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t *Transaction) IsWithDrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t *Transaction) ToResponseDTO() *dto.TransactionRes {
	return &dto.TransactionRes{
		TransactionID:   t.TransactionID,
		AccountID:       t.AccountID,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
