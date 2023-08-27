package domain

type CustomerRepositoryStub struct {
	customers []*Customer
}

func (r CustomerRepositoryStub) FindAll() ([]*Customer, error) {
	return r.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []*Customer{
		{ID: "1001", Name: "Proshanto", City: "Dhaka", Zipcode: "1212", DateOfBirth: "08/11/1997", Status: "active"},
		{ID: "1002", Name: "Oshanto", City: "Barishal", Zipcode: "1111", DateOfBirth: "12/01/1998", Status: "inactive"},
		{ID: "1003", Name: "Lol", City: "Rajshahi", Zipcode: "6969", DateOfBirth: "20/07/1999", Status: "active"},
	}

	return CustomerRepositoryStub{customers: customers}
}