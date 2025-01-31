package domain

import "bankapp/errs"

type Customer struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
	Dob     string `json:"dob"`
	Status  string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(string) (*Customer, *errs.AppError)
}
