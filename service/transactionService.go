package service

import (
	"time"

	"microservices.com/domain"
	"microservices.com/dto"
	"microservices.com/errors"
)

// TransactionService defines methods in the transaction service object
type TransactionService interface {
	NewTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError)
}

// DefaultTransactionService implements TransactionService
type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

// NewTransaction implements deposit method of TransactionService
func (s DefaultTransactionService) NewTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	transaction := domain.Transaction{
		AccountID:       req.AccountID,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, err := s.repo.Transact(&transaction)
	if err != nil {
		return nil, err
	}

	transactionRes := dto.TransactionResponse{
		TransactionID:  newTransaction.TransactionID,
		UpdatedBalance: newTransaction.Balance,
	}

	return &transactionRes, nil
}

// NewTransactionService creates the default service and injects repository
func NewTransactionService(repository domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repository}
}
