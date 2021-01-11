package dto

import "microservices.com/errors"

// TransactionRequest defines the shape of the transaction request body
type TransactionRequest struct {
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

// Validate will validate the request object
func (req *TransactionRequest) Validate() *errors.AppError {
	if req.TransactionType != "withdraw" && req.TransactionType != "deposit" {
		return errors.NewValidationError("Transaction type should be either withdraw or deposit")
	}

	if req.Amount < 0 {
		return errors.NewValidationError("Transaction amount cannot be negative")
	}

	return nil
}
