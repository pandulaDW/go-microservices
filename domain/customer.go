package domain

import "microservices.com/errors"

// Customer struct
type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     int
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

// CustomerRepository defines a repository
type CustomerRepository interface {
	FindAll(string) ([]Customer, *errors.AppError)
	ByID(string) (*Customer, *errors.AppError)
}
