package cmd

import (
	"encoding/json"
	"net/http"

	"github.com/H-proshanto/go-banking-microservice/banking/dto"
	"github.com/H-proshanto/go-banking-microservice/banking/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	svc *service.DefaultAccountSvc
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	customerID := queryParams["customer_id"]

	var req dto.NewAccountReq
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		req.CustomerID = customerID
		account, err := h.svc.NewAccount(&req)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	customerID := queryParams["customer_id"]
	accountID := queryParams["account_id"]

	var req dto.TransactionReq
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		req.AccountID = accountID
		req.CustomerID = customerID

		account, err := h.svc.MakeTransaction(&req)

		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}
