package errors

import "net/http"

// AppError defines the structure for application errors
type AppError struct {
	Code    int
	Message string
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
