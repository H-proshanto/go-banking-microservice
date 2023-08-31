package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/H-proshanto/go-banking-microservice/banking/domain"
	"github.com/H-proshanto/go-banking-microservice/banking/logger"
	"github.com/H-proshanto/go-banking-microservice/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	router := mux.NewRouter()

	dbClient := getDBClient()
	customerRepoDB := domain.NewCustomerRepoDB(dbClient)
	// accountRepoDB := domain.NewAccountRepoDB(dbClient)

	customerService := service.NewCustomerService(customerRepoDB)
	ch := CustomerHandlers{service: customerService}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	logger.Info("Server started on 8000")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}

func getDBClient() *sqlx.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", "localhost", "postgres", "password", "banking", "5432")
	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		logger.Error("Error connecting to DB client")
		panic(err)
	}

	return db
}
