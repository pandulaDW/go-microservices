package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"microservices.com/errors"
	"microservices.com/logger"
)

// TransactionRepoDB implements TransactionRepo
type TransactionRepoDB struct {
	client *sqlx.DB
}

// // Withdrawal will withdraw the given amount from the DB account
// func (t TransactionRepoDB) Withdrawal(amount float64) (Transaction, *errors.AppError) {

// }

// Deposit will withdraw the given amount from the DB account
func (r *TransactionRepoDB) Deposit(t *Transaction) (*Transaction, *errors.AppError) {
	sqlInsert := `INSERT INTO transactions 
	(account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)`

	result, err := r.client.Exec(sqlInsert, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	t.TransactionID = strconv.FormatInt(id, 10)
	return t, nil
}
