package domain

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver package
	"github.com/jmoiron/sqlx"
	"microservices.com/errors"
	"microservices.com/logger"
)

// CustomerRepositoryDB is the production repository, which contains
// a mysql database
type CustomerRepositoryDB struct {
	client *sqlx.DB
}

// FindAll returns all customers
func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errors.AppError) {
	findAllSQL := `SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers`
	findAllWithStatus := findAllSQL + ` WHERE status = ?`

	var err error
	customers := make([]Customer, 0)

	if status == "" {
		err = d.client.Select(&customers, findAllSQL)
	} else {
		err = d.client.Select(&customers, findAllWithStatus, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

// ByID returns a customer based on id
func (d CustomerRepositoryDB) ByID(id string) (*Customer, *errors.AppError) {
	customerSQL := `SELECT customer_id, name, city, zipcode, date_of_birth, status
				   FROM customers WHERE customer_id = ?`
	var c Customer
	err := d.client.Get(&c, customerSQL, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer table " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	return &c, nil
}

// NewCustomerRepositoryDb is a factory function for the CustomerRepositoryDB
func NewCustomerRepositoryDb() CustomerRepositoryDB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	client, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPassword, dbAddress, dbPort, dbName))
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{client}
}
