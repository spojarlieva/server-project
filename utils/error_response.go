package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse struct is standardized return for errors.
type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

// NewErrorResponse will create a new [ErrorResponse]
func NewErrorResponse(message string, statusCode int) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}

// RespondWithError will respond with [ErrorResponse].
// When passing value the pointer must not be nil.
func RespondWithError(w http.ResponseWriter, response *ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "There was an unknown error", http.StatusInternalServerError)
	}
}

// InternalServerError is the standardized return for server error.
func InternalServerError() *ErrorResponse {
	return NewErrorResponse(
		"Internal Server Error",
		http.StatusInternalServerError,
	)
}

// InvalidJson is the standardized return for invalid json that doesn't match a model.
func InvalidJson() *ErrorResponse {
	return NewErrorResponse(
		"Invalid Json",
		http.StatusBadRequest,
	)
}
