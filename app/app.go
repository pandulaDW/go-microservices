package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start registers the handler functions and starts the server
func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		log.Fatal(err)
	}
}
