package service

import "github.com/H-proshanto/go-banking-microservice/banking/domain"

type CustomerService interface {
	GetAllCustomers() ([]*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]*domain.Customer, error) {
	 return s.repo.FindAll()
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: repo,
	}
}