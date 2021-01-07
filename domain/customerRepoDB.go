package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver package
)

// CustomerRepositoryDB is the production repository, which contains
// a mysql database
type CustomerRepositoryDB struct {
	client *sql.DB
}

// FindAll returns all customers
func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err := d.client.Query(findAllSQL)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while querying customer table " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
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
