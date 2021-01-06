package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start registers the handler functions and starts the server
func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		log.Fatal(err)
	}
}
