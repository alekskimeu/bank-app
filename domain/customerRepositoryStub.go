package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Alex", City: "Nairobi", Zipcode: "00168", Dob: "1996-07-30", Status: "Active"},
		{Id: "1002", Name: "Jane", City: "Kiambu", Zipcode: "00200", Dob: "2002-01-20", Status: "Active"},
	}

	return CustomerRepositoryStub{customers}
}
