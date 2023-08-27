package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name string 	`json:"full_name" xml:"name"`
	City string 	`json:"city" xml:"city"`
	Zipcode string 	`json:"zip_code" xml:"zipcode"`
}

func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/customers", getAllCustomers)

	fmt.Println("Server started")
	http.ListenAndServe("0.0.0.0:8000", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")	
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []*Customer {
		{Name: "Pro", City: "Dhaka", Zipcode: "1212"},
		{Name: "sha", City: "Barishal", Zipcode: "1111"},
		{Name: "nto", City: "Rajshahi", Zipcode: "6969"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
