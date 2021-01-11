package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
	"microservices.com/errors"
	"microservices.com/logger"
)

// TransactionRepoDB implements TransactionRepo
type TransactionRepoDB struct {
	client *sqlx.DB
}

func getAccountBalance(tx *sqlx.Tx, id string) (*float64, *errors.AppError) {
	accountSQL := `SELECT account_id, amount FROM accounts WHERE account_id = ?`
	var a Account
	err := tx.Get(&a, accountSQL, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Account not found")
		}
		logger.Error("Error while scanning account table " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	return &a.Amount, nil
}

func commitTransaction(tx *sqlx.Tx, t *Transaction) *errors.AppError {
	sqlInsert := `INSERT INTO transactions 
	(account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)`

	result, err := tx.Exec(sqlInsert, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return errors.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return errors.NewUnexpectedError("Unexpected error from database")
	}

	t.TransactionID = strconv.FormatInt(id, 10)
	return nil
}

func updateAccount(tx *sqlx.Tx, t *Transaction) *errors.AppError {
	if t.TransactionType == "withdraw" && t.Balance < t.Amount {
		return errors.NewValidationError("Withdrawal amount cannot be larger than the balance")
	}

	sqlUpdate := `UPDATE accounts SET amount = ? WHERE account_id = ?`

	_, err := tx.Exec(sqlUpdate, t.Balance, t.AccountID)
	if err != nil {
		logger.Error("Error while updating the account: " + err.Error())
		return errors.NewUnexpectedError("Unexpected error from database")
	}

	return nil
}

// Transact will withdraw or deposit the given amount from the DB account
func (r TransactionRepoDB) Transact(t *Transaction) (*Transaction, *errors.AppError) {
	tx, err := r.client.Beginx()

	if err != nil {
		logger.Error("Error while initializing a transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	balance, appError := getAccountBalance(tx, t.AccountID)
	if appError != nil {
		return nil, appError
	}

	appError = commitTransaction(tx, t)
	if appError != nil {
		return nil, appError
	}

	if t.TransactionType == "withdraw" {
		t.Balance = *balance - t.Amount
	} else {
		t.Balance = *balance + t.Amount
	}

	appError = updateAccount(tx, t)
	if appError != nil {
		return nil, appError
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("Error while committing the transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	return t, nil
}

// NewTransactionRepoDB will create a new TransactionRepoDB
func NewTransactionRepoDB(client *sqlx.DB) TransactionRepoDB {
	return TransactionRepoDB{client: client}
}
