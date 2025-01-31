package domain

import (
	"bankapp/dto"
	"bankapp/errs"
)

type Customer struct {
	Id      string `db:"customer_id"`
	Name    string
	City    string
	Zipcode string
	Dob     string `db:"date_of_birth"`
	Status  string
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}

func (customer Customer) statusAsText() string {
	statusAsText := "actuve"

	if customer.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (customer Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:      customer.Id,
		Name:    customer.Name,
		City:    customer.City,
		Zipcode: customer.Zipcode,
		Dob:     customer.Dob,
		Status:  customer.statusAsText(),
	}
}
