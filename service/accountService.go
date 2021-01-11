package service

import (
	"time"

	"microservices.com/domain"
	"microservices.com/dto"
	"microservices.com/errors"
)

// AccountService interface implements a new account service
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError)
}

// DefaultAccountService defines the default account service
type DefaultAccountService struct {
	repo domain.AccountRepository
}

// NewAccount implements account service method
func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		CustomerID:  req.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	return newAccount.ToNewAccountResponseDTO(), nil
}

// NewAccountService creates a new DefaultAccountService
func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
