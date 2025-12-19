package response

import (
	"fmt"
	"net/http"
)

// APIError represents a structured error for API responses
type APIError struct {
	Code    int         `json:"code"`            // HTTP status code
	Message string      `json:"message"`         // Error message
	Err     interface{} `json:"error,omitempty"` // Underlying error details
}

func (e *APIError) Error() string {
	switch v := e.Err.(type) {
	case string:
		return v
	case error:
		return v.Error()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func NewAPIError(statusCode int, message string, err interface{}) *APIError {
	return &APIError{
		Code:    statusCode,
		Message: message,
		Err:     err,
	}
}

func NewBadRequestError(message string, err interface{}) *APIError {
	return NewAPIError(http.StatusBadRequest, message, err)
}

func NewNotFoundError(message string, err interface{}) *APIError {
	return NewAPIError(http.StatusNotFound, message, err)
}

func NewInternalServerError(message string, err interface{}) *APIError {
	return NewAPIError(http.StatusInternalServerError, message, err)
}
