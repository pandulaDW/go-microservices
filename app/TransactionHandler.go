package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"microservices.com/dto"
	"microservices.com/service"
)

// TransactionHandler creates a handler object with service as a dependency
type TransactionHandler struct {
	service service.TransactionService
}

// Deposit will deposit a new transaction
func (h TransactionHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	var request dto.TransactionRequest
	vars := mux.Vars(r)
	request.AccountID = vars["account_id"]
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	transaction, appError := h.service.NewDeposit(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, transaction)
}
