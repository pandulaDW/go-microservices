package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"microservices.com/dto"
	"microservices.com/service"
)

// AccountHandler injects an account service into account handler
type AccountHandler struct {
	service service.AccountService
}

// NewAccount handler creates a new account
func (h *AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	vars := mux.Vars(r)
	request.CustomerID = vars["customer_id"]
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	account, appError := h.service.NewAccount(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, account)
}
