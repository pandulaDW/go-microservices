package service

import (
	"microservices.com/domain"
	"microservices.com/dto"
	"microservices.com/errors"
)

// CustomerService defines a service
type CustomerService interface {
	GetAllCustomers(string) ([]*dto.CustomerResponse, *errors.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errors.AppError)
}

// DefaultCustomerService is the implementation of the service
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomers returns all customers
func (s DefaultCustomerService) GetAllCustomers(status string) ([]*dto.CustomerResponse, *errors.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	}
	customers, err := s.repo.FindAll(status)
	customerDTOs := make([]*dto.CustomerResponse, 0, len(customers))
	for _, val := range customers {
		customerDTOs = append(customerDTOs, val.ToDto())
	}
	return customerDTOs, err
}

// GetCustomer returns a customer by id
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errors.AppError) {
	c, err := s.repo.ByID(id)
	if err != nil {
		return nil, err
	}
	return c.ToDto(), nil
}

// NewCustomerService creates the default service and injects repository
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
