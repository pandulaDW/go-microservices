package domain

import (
	"microservices.com/dto"
	"microservices.com/errors"
)

// Account defines the account db object
type Account struct {
	AccountID   string `db:"account_id"`
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64 `db:"amount"`
	Status      string
}

// ToNewAccountResponseDTO converts the account object to the response dto
func (a Account) ToNewAccountResponseDTO() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountID: a.AccountID}
}

// AccountRepository interface defines account repo methods
type AccountRepository interface {
	Save(Account) (*Account, *errors.AppError)
}
