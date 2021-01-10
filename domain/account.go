package domain

import "microservices.com/errors"

// Account defines the account db object
type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

// AccountRepository interface defines account repo methods
type AccountRepository interface {
	Save(Account) (*Account, *errors.AppError)
}
