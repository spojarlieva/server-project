package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"server/auth"
	"server/services"
	"server/utils"
	"strconv"
)

// EventHandler will respond to request related to events.
type EventHandler interface {
	// GetEvents will return all events.
	GetEvents() http.HandlerFunc

	// RegisterForEvents will register a user for an event.
	RegisterForEvents() http.HandlerFunc

	// UnregisterForEvents will remove registration for an event.
	UnregisterForEvents() http.HandlerFunc

	// GetRegisteredEvents will return events with additional field if the user
	// has registered for them.
	GetRegisteredEvents() http.HandlerFunc
}

// DefaultEventHandler is default implementation of [EventHandler].
type DefaultEventHandler struct {
	eventRepository services.EventService
}

func (h *DefaultEventHandler) GetEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, errorResponse := h.eventRepository.GetEvents(r.Context())
		if errorResponse != nil {
			utils.RespondWithError(w, errorResponse)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(events)
		if err != nil {
			utils.RespondWithError(w, utils.InternalServerError())
		}
	}
}

func (h *DefaultEventHandler) RegisterForEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(auth.Key).(*jwt.RegisteredClaims)
		if !ok {
			utils.RespondWithError(w, utils.InternalServerError())
			return
		}
		userId, err := strconv.Atoi(claims.Subject)
		if err != nil {
			utils.RespondWithError(w, utils.InternalServerError())
			return
		}

		id := r.PathValue("id")
		eventId, err := strconv.Atoi(id)
		if err != nil {
			utils.RespondWithError(w, utils.NewErrorResponse("Invalid id", http.StatusBadRequest))
			return
		}

		errorResponse := h.eventRepository.RegisterForEvent(r.Context(), userId, eventId)
		if errorResponse != nil {
			utils.RespondWithError(w, errorResponse)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *DefaultEventHandler) UnregisterForEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(auth.Key).(*jwt.RegisteredClaims)
		if !ok {
			utils.RespondWithError(w, utils.InternalServerError())
			return
		}
		userId, err := strconv.Atoi(claims.Subject)
		if err != nil {
			utils.RespondWithError(w, utils.InternalServerError())
			return
		}

		id := r.PathValue("id")
		eventId, err := strconv.Atoi(id)
		if err != nil {
			utils.RespondWithError(w, utils.NewErrorResponse("Invalid id", http.StatusBadRequest))
			return
		}

		errorResponse := h.eventRepository.UnregisterForEvent(r.Context(), userId, eventId)
		if errorResponse != nil {
			utils.RespondWithError(w, errorResponse)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *DefaultEventHandler) GetRegisteredEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(auth.Key).(*jwt.RegisteredClaims)
		if !ok {
			utils.RespondWithError(w, utils.InternalServerError())
			return
		}
		userId, err := strconv.Atoi(claims.Subject)
		if err != nil {
			utils.RespondWithError(w, utils.InternalServerError())
			return
		}

		events, errorResponse := h.eventRepository.GetRegisteredEvents(r.Context(), userId)
		if errorResponse != nil {
			utils.RespondWithError(w, errorResponse)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(events)
		if err != nil {
			utils.RespondWithError(w, utils.InternalServerError())
		}
	}
}

// NewDefaultEventHandler will create new [DefaultEventHandler].
func NewDefaultEventHandler(eventRepository services.EventService) *DefaultEventHandler {
	return &DefaultEventHandler{
		eventRepository: eventRepository,
	}
}
