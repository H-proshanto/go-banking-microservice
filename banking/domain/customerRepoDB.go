package domain

import (
	"database/sql"
	"fmt"

	"github.com/H-proshanto/go-banking-microservice/banking/errs"
	"github.com/H-proshanto/go-banking-microservice/banking/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepoDB struct {
	db *sqlx.DB
}

func NewCustomerRepoDB() *CustomerRepoDB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", "localhost", "postgres", "password", "banking", "5432")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	return &CustomerRepoDB{
		db: db,
	}

}

func (r *CustomerRepoDB) FindAll(status string) ([]*Customer, *errs.AppError) {
	var err error
	customers := make([]*Customer, 0)

	if status == "" {
		findAllSqlQuery := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = r.db.Select(&customers, findAllSqlQuery)
	} else {
		logger.Error(status)
		findAllSqlQuery := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers WHERE status = $1"
		err = r.db.Select(&customers, findAllSqlQuery, status)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return customers, nil
}

func (r *CustomerRepoDB) ById(id string) (*Customer, *errs.AppError) {
	findCustomerByIdSqlQuery := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = $1"

	var c Customer
	err := r.db.Get(&c, findCustomerByIdSqlQuery, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil

}
