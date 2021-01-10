package domain

import (
	"microservices.com/dto"
	"microservices.com/errors"
)

// Customer struct
type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     int
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c *Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

// ToDto transforms the domain object to dto
func (c *Customer) ToDto() *dto.CustomerResponse {
	response := dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
	return &response
}

// CustomerRepository defines a repository
type CustomerRepository interface {
	FindAll(string) ([]Customer, *errors.AppError)
	ByID(string) (*Customer, *errors.AppError)
}
