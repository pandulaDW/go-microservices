package domain

// CustomerRepositoryStub is a stub
type CustomerRepositoryStub struct {
	customers []Customer
}

// FindAll returns all customers
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// NewCustomerRepositoryStub is a factory function for CustomerRepositoryStub
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Jack", "London", 110075, "2000-01-29", "1"},
		{"1002", "Rob", "Canada", 120035, "2004-01-29", "1"},
	}

	return CustomerRepositoryStub{customers}
}
