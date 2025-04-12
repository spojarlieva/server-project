package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/services"
	"server/utils"
)

// UserHandler interface will respond to request related to users.
type UserHandler interface {
	// Register method will register a user.
	Register() http.HandlerFunc

	// Login create user session with jwt.
	Login() http.HandlerFunc
}

// DefaultUserHandler struct is default implementation of [UserHandler].
type DefaultUserHandler struct {
	userService services.UserService
}

func (h *DefaultUserHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userPayload models.UserPayload

		err := json.NewDecoder(r.Body).Decode(&userPayload)
		if err != nil {
			utils.RespondWithError(w, utils.InvalidJson())
			return
		}

		if returned := utils.CheckPayload(w, &userPayload); returned {
			return
		}

		errorResponse := h.userService.AddUser(r.Context(), &userPayload)
		if errorResponse != nil {
			utils.RespondWithError(w, errorResponse)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (h *DefaultUserHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userPayload models.UserPayload

		err := json.NewDecoder(r.Body).Decode(&userPayload)
		if err != nil {
			utils.RespondWithError(w, utils.InvalidJson())
			return
		}

		token, errorResponse := h.userService.Login(r.Context(), &userPayload)
		if errorResponse != nil {
			utils.RespondWithError(w, errorResponse)
			return
		}

		response := map[string]string{
			"token": token,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			utils.RespondWithError(w, utils.InternalServerError())
		}
	}
}

// NewDefaultUserHandler will create a new [DefaultUserHandler].
func NewDefaultUserHandler(userService services.UserService) *DefaultUserHandler {
	return &DefaultUserHandler{
		userService: userService,
	}
}
