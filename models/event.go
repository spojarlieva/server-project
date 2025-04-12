package models

import "time"

// Event struct holds event information.
type Event struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Address     string    `json:"address"`
	ImageURL    string    `json:"image_url"`
	Description string    `json:"description"`
}

// EventWithRegistration struct holds event with information plus
// information if the user has registered for it.
type EventWithRegistration struct {
	Event
	IsRegistered bool `json:"is_registered"`
}
