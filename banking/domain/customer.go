package domain

import "github.com/H-proshanto/go-banking-microservice/banking/errs"

type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface { 
	FindAll(status string) ([]*Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError) 
}
