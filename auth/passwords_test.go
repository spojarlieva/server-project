package auth

import (
	"testing"
)

func TestHashAndVerifyPasswordCorrect(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	err = VerifyPassword(hash, password)
	if err != nil {
		t.Errorf("Error verifying password: %v", err)
	}
}

func TestHashAndVerifyPasswordIncorrect(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	err = VerifyPassword(hash, "other password")
	if err == nil {
		t.Errorf("Error verifying password: %v", err)
	}
}
