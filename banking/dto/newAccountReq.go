package dto

import "github.com/H-proshanto/go-banking-microservice/banking/errs"

type NewAccountReq struct {
	CustomerID  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r *NewAccountReq) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need at least 5000")
	}
	if r.AccountType != "saving" && r.AccountType != "checking" {
		return errs.NewValidationError("Account type should be either checking or saving")
	}

	return nil
}
