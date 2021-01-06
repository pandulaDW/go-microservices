package domain

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
	FindAll() ([]Customer, error)
}
