package domain

import (
	"log"

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

func (r *CustomerRepoDB) FindAll() ([]*Customer, error) {

	findAllSqlQuery := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := r.db.Raw(findAllSqlQuery).Rows()

	if err != nil {
		log.Println("Error while querying customer table" + err.Error())

		return nil, err
	}

	defer rows.Close()

	customers := make([]*Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

		if err != nil {
			log.Println("Error while scanning customers" + err.Error())

			return nil, err
		}
		customers = append(customers, &c)
	}

	return customers, nil

}
