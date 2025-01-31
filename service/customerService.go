package service

import (
	"bankapp/domain"
	"bankapp/dto"
	"bankapp/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {

	customer, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	response := dto.CustomerResponse{
		Id:      customer.Id,
		Name:    customer.Name,
		City:    customer.City,
		Zipcode: customer.Zipcode,
		Dob:     customer.Dob,
		Status:  customer.Status,
	}
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
