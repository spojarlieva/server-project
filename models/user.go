package models

import (
	"net/http"
	"server/utils"
	"strings"
)

// UserPayload holds used information.
type UserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserPayload) Validate() *utils.ErrorResponse {
	result := u.CheckEmail()
	if result != nil {
		return result
	}

	result = u.CheckPassword()
	return result
}

func (u *UserPayload) CheckEmail() *utils.ErrorResponse {
	if u.Email == "" {
		return utils.NewErrorResponse("Email is empty", http.StatusBadRequest)
	}

	splitEmail := strings.Split(u.Email, "@")
	if len(splitEmail) != 2 {
		return utils.NewErrorResponse("Email is invalid", http.StatusBadRequest)
	}

	if splitEmail[0] == "" {
		return utils.NewErrorResponse("Missing local part of the email", http.StatusBadRequest)
	}

	if splitEmail[1] == "" {
		return utils.NewErrorResponse("Missing domain part of the email", http.StatusBadRequest)
	}

	splitDomain := strings.Split(splitEmail[1], ".")
	if len(splitDomain) != 2 {
		return utils.NewErrorResponse("Domain is invalid", http.StatusBadRequest)
	}

	if len(splitDomain[0]) < 1 || len(splitDomain[1]) < 2 {
		return utils.NewErrorResponse("Domain is invalid", http.StatusBadRequest)
	}
	return nil
}

func (u *UserPayload) CheckPassword() *utils.ErrorResponse {
	if u.Password == "" {
		return utils.NewErrorResponse("Password is empty", http.StatusBadRequest)
	}

	if !strings.ContainsAny(u.Password, "0123456789") {
		return utils.NewErrorResponse("Password must contain a number", http.StatusBadRequest)
	}

	if !strings.ContainsAny(u.Password, "abcdefghijklmnopqrstuvwxyz") {
		return utils.NewErrorResponse("Password must contain a letter", http.StatusBadRequest)
	}

	if !strings.ContainsAny(u.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return utils.NewErrorResponse("Password must contain a capital letter", http.StatusBadRequest)
	}

	if !strings.ContainsAny(u.Password, "!@#$^&*()_-+") {
		return utils.NewErrorResponse("Password must contain a special character", http.StatusBadRequest)
	}

	return nil
}

// User struct holds used data locally.
type User struct {
	Id       int
	Email    string
	Password string
}
