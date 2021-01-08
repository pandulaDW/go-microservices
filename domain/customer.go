package domain

import "microservices.com/errors"

// Customer struct
type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     int
	DateofBirth string
	Status      string
}

// CustomerRepository defines a repository
type CustomerRepository interface {
	FindAll() ([]Customer, *errors.AppError)
	ByID(string) (*Customer, *errors.AppError)
}
