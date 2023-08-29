package domain

import (
	"database/sql"
	"log"

	"github.com/H-proshanto/go-banking-microservice/banking/errs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerRepoDB struct {
	db *gorm.DB
}

func NewCustomerRepoDB() *CustomerRepoDB {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &CustomerRepoDB{
		db: db,
	}

}

func (r *CustomerRepoDB) FindAll() ([]*Customer, *errs.AppError) {

	findAllSqlQuery := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := r.db.Raw(findAllSqlQuery).Rows()

	if err != nil {
		log.Println("Error while querying customer table" + err.Error())

		return nil, errs.NewUnexpectedError("something unexpected happend")
	}

	defer rows.Close()

	customers := make([]*Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("Customers not found")
			} else {
				log.Println("Error while scanning customers" + err.Error())
				return nil, errs.NewUnexpectedError("unexpected database error")
			}
		}
		customers = append(customers, &c)
	}

	return customers, nil

}

func (r *CustomerRepoDB) ById(id string) (*Customer, *errs.AppError) {
	findCustomerByIdSqlQuery := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"

	row := r.db.Raw(findCustomerByIdSqlQuery, id).Row()

	var c Customer

	err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil

}
