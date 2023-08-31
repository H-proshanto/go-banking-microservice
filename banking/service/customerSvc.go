package service

import (
	"github.com/H-proshanto/go-banking-microservice/banking/domain"
	"github.com/H-proshanto/go-banking-microservice/banking/dto"
	"github.com/H-proshanto/go-banking-microservice/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]*dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]*dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := s.repo.FindAll(status)

	if err != nil {
		return nil, err
	}

	var response []*dto.CustomerResponse
	for _, customer := range customers {
		response = append(response, customer.ToResponseDTO())
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	return customer.ToResponseDTO(), nil
}

func NewCustomerService(repo domain.CustomerRepository) *DefaultCustomerService {
	return &DefaultCustomerService{
		repo: repo,
	}
}
