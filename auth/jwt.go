package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"server/utils"
	"strconv"
	"strings"
	"time"
)

// TokenKey type used to store tokens in [context.Context]
type TokenKey string

// Key is variable used to get the token from [context.Context]
const Key TokenKey = "token"

// JWTAuthenticator struct will handle authentication of the api.
type JWTAuthenticator struct {
	// secret used to hash tokens.
	secret []byte
	// issuer of the tokens.
	issuer string
}

// CreateToken method will create a new token and set the subject as used id.
func (a *JWTAuthenticator) CreateToken(userId int) (string, error) {
	token := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userId),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	hashToken := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	tokenString, err := hashToken.SignedString(a.secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// validateToken method will parse a token into [jwt.RegisteredClaims].a
func (a *JWTAuthenticator) validateToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// Middleware method will add a middleware to [http.HandleFunc] by checking the token
// and passing it to the next handler. If the token is invalid it will drop the request.
func (a *JWTAuthenticator) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := a.validateToken(tokenString)
		if err != nil {
			utils.RespondWithError(w, utils.NewErrorResponse("Invalid Token", http.StatusUnauthorized))
			return
		}

		ctx := context.WithValue(r.Context(), Key, claims)
		next(w, r.WithContext(ctx))
	}
}

// NewJWTAuthenticator will create a new [JWTAuthenticator].
func NewJWTAuthenticator(secret []byte, issuer string) *JWTAuthenticator {
	return &JWTAuthenticator{
		secret,
		issuer,
	}
}
