package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/H-proshanto/go-banking-microservice/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()

	if err != nil {
		fmt.Fprint(w, "could not get customers")
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
