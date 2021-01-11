package domain

import "microservices.com/errors"

// Transaction defines the shape of a transaction
type Transaction struct {
	TransactionID   string
	AccountID       string
	Amount          float64
	Balance         float64
	TransactionType string
	TransactionDate string
}

// TransactionRepository interface defines transaction methods
type TransactionRepository interface {
	// Withdrawal(amount float64) (*Transaction, *errors.AppError)
	Deposit(t *Transaction) (*Transaction, *errors.AppError)
}
