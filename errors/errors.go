package errors

import "net/http"

// AppError defines the structure for application errors
type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

// AsMessage will return the error with the code
func (e *AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

// NewNotFoundError returns a not found error
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

// NewUnexpectedError returns a not found error
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
