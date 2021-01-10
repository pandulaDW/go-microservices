// dto deals with the user side

package dto

// CustomerResponse defines the user side transforamtion of customer
type CustomerResponse struct {
	ID          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     int    `json:"zipcode"`
	DateofBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}
