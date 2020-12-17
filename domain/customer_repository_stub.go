package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "John", City: "Night City", DateOfBirth: "12/12/83", Zipcode: "123", Status: "1"},
		{Id: "2", Name: "Maria", City: "Night City", DateOfBirth: "12/12/83", Zipcode: "123", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
