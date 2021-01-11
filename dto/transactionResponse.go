package dto

// TransactionResponse defines the response body for a new transaction request
type TransactionResponse struct {
	TransactionID  string  `json:"transaction_id"`
	UpdatedBalance float64 `json:"updated_balance"`
}
