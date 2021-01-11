package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"microservices.com/domain"
	"microservices.com/service"
)

func getDbClient() *sqlx.DB {
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

	return client
}

// Start registers the handler functions and starts the server
func Start() {
	router := mux.NewRouter()
	dbClient := getDbClient()

	// create repositories
	// stubRepo := domain.NewCustomerRepositoryStub()
	dbCustomerRepo := domain.NewCustomerRepositoryDb(dbClient)
	dbAccountRepo := domain.NewAccountRepoDB(dbClient)

	// create services
	customerService := service.NewCustomerService(dbCustomerRepo)
	accountService := service.NewAccountService(dbAccountRepo)

	// create handlers
	customerHandler := CustomerHandler{customerService}
	accountHandler := AccountHandler{accountService}

	// define routes
	router.HandleFunc("/customers", customerHandler.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", customerHandler.getAllCustomers).Methods(http.MethodGet).Queries("{status:active}", "{status:inactive}")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", accountHandler.NewAccount).Methods(http.MethodPost)

	serverAddress := fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), os.Getenv("PORT"))

	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Fatal(err)
	}
}
