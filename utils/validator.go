package utils

import "net/http"

// ValidatablePayload interface used to validate payload.
type ValidatablePayload interface {
	// Validate method used to check if the payload is ok.
	// Return [ErrorResponse] if there is an error or nil if everything is correct.
	Validate() *ErrorResponse
}

// CheckPayload will check the payload and respond with proper error if there is any.
// Returns true if there was an error, and it responded otherwise it returns false.
func CheckPayload(w http.ResponseWriter, payload ValidatablePayload) bool {
	errorResponse := payload.Validate()
	if errorResponse != nil {
		RespondWithError(w, errorResponse)
		return true
	}
	return false
}
