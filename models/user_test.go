package models

import (
	"net/http"
	"reflect"
	"server/utils"
	"testing"
)

type testCase struct {
	name     string
	user     UserPayload
	expected *utils.ErrorResponse
}

func TestUserValidate(t *testing.T) {
	cases := []testCase{
		{
			name: "Valid user",
			user: UserPayload{
				Email:    "user@example.com",
				Password: "Strong_Password1",
			},
			expected: nil,
		}, {
			name: "Invalid email(empty)",
			user: UserPayload{
				Email:    "",
				Password: "Strong_Password1",
			},
			expected: utils.NewErrorResponse("Email is empty", http.StatusBadRequest),
		}, {
			name: "Invalid email(missing @)",
			user: UserPayload{
				Email:    "usereample.com",
				Password: "Strong_Password1",
			},
			expected: utils.NewErrorResponse("Email is invalid", http.StatusBadRequest),
		}, {
			name: "Invalid email(missing local part)",
			user: UserPayload{
				Email:    "@example.com",
				Password: "Strong_Password1",
			},
			expected: utils.NewErrorResponse("Missing local part of the email", http.StatusBadRequest),
		}, {
			name: "Invalid email(missing domain part)",
			user: UserPayload{
				Email:    "user@",
				Password: "Strong_Password1",
			},
			expected: utils.NewErrorResponse("Missing domain part of the email", http.StatusBadRequest),
		}, {
			name: "Invalid email(invalid domain part)",
			user: UserPayload{
				Email:    "user@e.c",
				Password: "Strong_Password1",
			},
			expected: utils.NewErrorResponse("Domain is invalid", http.StatusBadRequest),
		}, {
			name: "Invalid password(empty)",
			user: UserPayload{
				Email:    "user@example.com",
				Password: "",
			},
			expected: utils.NewErrorResponse("Password is empty", http.StatusBadRequest),
		}, {
			name: "Invalid password(missing number)",
			user: UserPayload{
				Email:    "user@example.com",
				Password: "Strong_Password",
			},
			expected: utils.NewErrorResponse("Password must contain a number", http.StatusBadRequest),
		}, {
			name: "Invalid password(missing letter)",
			user: UserPayload{
				Email:    "user@example.com",
				Password: "_1234",
			},
			expected: utils.NewErrorResponse("Password must contain a letter", http.StatusBadRequest),
		}, {
			name: "Invalid password(missing capital letter)",
			user: UserPayload{
				Email:    "user@example.com",
				Password: "strong_password1",
			},
			expected: utils.NewErrorResponse("Password must contain a capital letter", http.StatusBadRequest),
		}, {
			name: "Invalid password(missing special character)",
			user: UserPayload{
				Email:    "user@example.com",
				Password: "StrongPassword1",
			},
			expected: utils.NewErrorResponse("Password must contain a special character", http.StatusBadRequest),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.user.Validate()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Validate() expected %v, got %v", tc.expected, result)
			}
		})
	}
}
