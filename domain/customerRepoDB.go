package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver package
	"microservices.com/errors"
)

// CustomerRepositoryDB is the production repository, which contains
// a mysql database
type CustomerRepositoryDB struct {
	client *sql.DB
}

// FindAll returns all customers
func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errors.AppError) {
	findAllSQL := `SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers`
	findAllWithStatus := findAllSQL + ` WHERE status = ?`

	var rows *sql.Rows
	var err error

	if status == "" {
		rows, err = d.client.Query(findAllSQL)
	} else {
		rows, err = d.client.Query(findAllWithStatus, status)
	}

	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customer table " + err.Error())
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

// ByID returns a customer based on id
func (d CustomerRepositoryDB) ByID(id string) (*Customer, *errors.AppError) {
	customerSQL := `SELECT customer_id, name, city, zipcode, date_of_birth, status
				   FROM customers WHERE customer_id = ?`

	row := d.client.QueryRow(customerSQL, id)
	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Customer not found")
		}
		log.Println("Error while scaning customer " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	return &c, nil
}

// NewCustomerRepositoryDb is a factory function for the CustomerRepositoryDB
func NewCustomerRepositoryDb() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:my-secret-pw@tcp(192.168.99.100:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{client}
}
