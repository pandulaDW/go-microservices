package app

import (
	"net/http"

	"microservices.com/service"
)

// TransactionHandler creates a handler object with service as a dependency
type TransactionHandler struct {
	service service.TransactionService
}

// Deposit will deposit a new transaction
func (h TransactionHandler) Deposit(w http.ResponseWriter, r *http.Request) {

}
