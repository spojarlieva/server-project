package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"net/http/httptest"
	"testing"
)

var authenticator = NewJWTAuthenticator([]byte("secret"), "issuer")

func TestJWTAuthenticatorCreateToken(t *testing.T) {
	_, err := authenticator.CreateToken(1)
	if err != nil {
		t.Fatalf("Error creating token: %v", err)
	}
}

func TestJWTAuthenticatorValidateToken(t *testing.T) {
	stringToken, err := authenticator.CreateToken(1)
	if err != nil {
		t.Fatalf("Error creating token: %v", err)
		return
	}

	parsedToken, err := authenticator.validateToken(stringToken)
	if err != nil {
		t.Fatalf("Error validating token: %v", err)
		return
	}

	if parsedToken == nil {
		t.Fatalf("Nil token")
		return
	}

	if parsedToken.Subject != "1" {
		t.Fatalf("Token id should be 1, but was %v", parsedToken.ID)
		return
	}
}

func TestJWTAuthenticatorMiddleware(t *testing.T) {
	token, err := authenticator.CreateToken(1)
	if err != nil {
		t.Fatalf("Error creating token: %v", err)
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(Key).(*jwt.RegisteredClaims)
		if !ok || claims.Subject != "1" {
			t.Fatalf("Passed invalid claims")
		}

		w.WriteHeader(http.StatusOK)
	})

	handler := authenticator.Middleware(nextHandler)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Authorization", "Bearer "+token)

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("Unexpected status code with valid token: %v", recorder.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Authorization", "invalid token ")
	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("Unexpected status code with invalid token: %v", recorder.Code)
	}
}
