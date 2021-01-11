package dto

import (
	"strings"

	"microservices.com/errors"
)

// NewAccountRequest defines the request body for an account request
type NewAccountRequest struct {
	CustomerID  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

// Validate will validate the request object
func (r *NewAccountRequest) Validate() *errors.AppError {
	if r.Amount < 5000 {
		return errors.NewValidationError("To open a new account you need to deposit at least 5000rs")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errors.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
