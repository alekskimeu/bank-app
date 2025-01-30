package domain

type Customer struct {
	Id      string
	Name    string
	City    string
	Zipcode string
	Dob     string
	Status  string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(string) (*Customer, error)
}
