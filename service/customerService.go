package service

import "microservices.com/domain"

// CustomerService defines a service
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

// DefaultCustomerService is the implementation of the service
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomers returns all customers
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// NewCustomerService creates the default service and injects repository
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
