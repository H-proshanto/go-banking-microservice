package domain

import (
	"github.com/H-proshanto/go-banking-microservice/banking/dto"
	"github.com/H-proshanto/go-banking-microservice/banking/errs"
)

type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c *Customer) statusAsText() string {
	statusAsText := "active"

	if c.Status == "0" {
		statusAsText = "inactive"
	}
	
	return statusAsText
}

func (c *Customer) ToResponseDTO() *dto.CustomerResponse {


	return &dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]*Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}
