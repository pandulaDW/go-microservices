package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"microservices.com/domain"
	"microservices.com/service"
)

// Start registers the handler functions and starts the server
func Start() {
	router := mux.NewRouter()

	// create customer repositories
	// stubRepo := domain.NewCustomerRepositoryStub()
	dbRepo := domain.NewCustomerRepositoryDb()

	// create customer service
	customerService := service.NewCustomerService(dbRepo)

	// create customer handler
	customerHandler := CustomerHandler{customerService}

	// define routes
	router.HandleFunc("/customers", customerHandler.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getCustomer).Methods(http.MethodGet)

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		log.Fatal(err)
	}
}
