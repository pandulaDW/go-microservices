package service

import (
	"microservices.com/domain"
	"microservices.com/errors"
)

// CustomerService defines a service
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errors.AppError)
}

// DefaultCustomerService is the implementation of the service
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomers returns all customers
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// GetCustomer returns a customer by id
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errors.AppError) {
	return s.repo.ByID(id)
}

// NewCustomerService creates the default service and injects repository
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
