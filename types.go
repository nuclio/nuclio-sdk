package nuclio

import (
	"github.com/satori/go.uuid"
)

// ID is an event ID
type ID *uuid.UUID

// NewID return a new random ID
func NewID() ID {
	// create a unique request ID
	id := uuid.NewV4()
	return ID(&id)
}

// EventHandler is function signature and event handler function
type EventHandler func(*Context, Event) (interface{}, error)
