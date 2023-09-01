package service

import (
	"time"

	"github.com/H-proshanto/go-banking-microservice/banking/domain"
	"github.com/H-proshanto/go-banking-microservice/banking/dto"
	"github.com/H-proshanto/go-banking-microservice/banking/errs"
)

type AccountService interface {
	NewAccount(req *dto.NewAccountReq) (*dto.NewAccountRes, *errs.AppError)
}

type DefaultAccountSvc struct {
	repo domain.AccountRepo
}

func (s DefaultAccountSvc) NewAccount(req *dto.NewAccountReq) (*dto.NewAccountRes, *errs.AppError) {
	err := req.Validate()

	if err != nil {
		return nil, err
	}

	account := domain.Account{
		AccountID:   "",
		CustomerID:  req.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := s.repo.Save(&account)

	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResDTO()

	return response, nil
}

func NewAccountSvc(repo domain.AccountRepo) *DefaultAccountSvc {
	return &DefaultAccountSvc{
		repo: repo,
	}
}

func (s DefaultAccountSvc) MakeTransaction(req *dto.TransactionReq) (*dto.TransactionRes, *errs.AppError) {
	err := req.Validate()

	if err != nil {
		return nil, err
	}

	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(&req.AccountID)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in account")
		}
	}

	t := domain.Transaction{
		AccountID:       req.AccountID,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	transaction, err := s.repo.SaveTransaction(&t)
	if err != nil {
		return nil, err
	}

	response := transaction.ToResponseDTO()
	return response, nil
}
