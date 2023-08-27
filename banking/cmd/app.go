package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()
	router.HandleFunc("/test", test).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id}", getCustomer).Methods(http.MethodGet)

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}