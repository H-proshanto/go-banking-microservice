package service

import (
	"github.com/H-proshanto/go-banking-microservice/banking/domain"
	"github.com/H-proshanto/go-banking-microservice/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]*domain.Customer, *errs.AppError)
	GetCustomer(id string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]*domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repo domain.CustomerRepository) *DefaultCustomerService {
	return &DefaultCustomerService{
		repo: repo,
	}
}
