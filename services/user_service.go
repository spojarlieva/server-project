package services

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"server/auth"
	"server/models"
	"server/repositories"
	"server/utils"
)

// UserService interface will manage users business logic.
type UserService interface {
	// AddUser will check the email, encrypt the password and add the user.
	AddUser(ctx context.Context, user *models.UserPayload) *utils.ErrorResponse

	// Login will check user credentials and if they are correct will return jwt.
	Login(ctx context.Context, user *models.UserPayload) (string, *utils.ErrorResponse)
}

// DefaultUserService struct is the default implementation of [UserService].
type DefaultUserService struct {
	userRepository repositories.UserRepository
	authenticator  *auth.JWTAuthenticator
}

func (s *DefaultUserService) AddUser(ctx context.Context, user *models.UserPayload) *utils.ErrorResponse {
	result, err := s.userRepository.CheckIfEmailExists(ctx, user.Email)
	if err != nil {
		return utils.InternalServerError()
	}

	if result {
		return utils.NewErrorResponse(
			"The email is already taken",
			http.StatusConflict,
		)
	}

	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		return utils.InternalServerError()
	}

	err = s.userRepository.AddUser(ctx, user)
	if err != nil {
		return utils.InternalServerError()
	}

	return nil
}

func (s *DefaultUserService) Login(ctx context.Context, user *models.UserPayload) (string, *utils.ErrorResponse) {
	fetchedUser, err := s.userRepository.GetUserByEmail(ctx, user.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return "", utils.NewErrorResponse("Invalid credentials", http.StatusUnauthorized)
	} else if err != nil {
		return "", utils.InternalServerError()
	}

	err = auth.VerifyPassword(fetchedUser.Password, user.Password)
	if err != nil {
		return "", utils.NewErrorResponse("Invalid credentials", http.StatusUnauthorized)
	}

	token, err := s.authenticator.CreateToken(fetchedUser.Id)
	if err != nil {
		return "", utils.InternalServerError()
	}

	return token, nil
}

// NewDefaultUserService will create a new [DefaultUserService].
func NewDefaultUserService(userRepository repositories.UserRepository, authenticator *auth.JWTAuthenticator) *DefaultUserService {
	return &DefaultUserService{
		userRepository: userRepository,
		authenticator:  authenticator,
	}
}
