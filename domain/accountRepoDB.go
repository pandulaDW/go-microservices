package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"microservices.com/errors"
	"microservices.com/logger"
)

// AccountRepositoryDB is an implementation of AccountRepository
type AccountRepositoryDB struct {
	client *sqlx.DB
}

// Save will save a new account in the db
func (d AccountRepositoryDB) Save(a Account) (*Account, *errors.AppError) {
	sqlInsert := `INSERT INTO accounts 
	(customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)`

	result, err := d.client.Exec(sqlInsert, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountID = strconv.FormatInt(id, 10)
	return &a, nil
}

// NewAccountRepoDB is a constructor function for AccountRepositoryDB
func NewAccountRepoDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
