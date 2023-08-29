package service

import (
	"github.com/H-proshanto/go-banking-microservice/banking/domain"
	"github.com/H-proshanto/go-banking-microservice/banking/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]*domain.Customer, *errs.AppError)
	GetCustomer(id string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]*domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repo domain.CustomerRepository) *DefaultCustomerService {
	return &DefaultCustomerService{
		repo: repo,
	}
}
