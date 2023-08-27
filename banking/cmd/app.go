package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/H-proshanto/go-banking-microservice/banking/domain"
	"github.com/H-proshanto/go-banking-microservice/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()


	ch := CustomerHandlers{service : service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	

	fmt.Printf("Server started on %v", "0.0.0.0:8000")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}