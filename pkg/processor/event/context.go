package event

import (
	"github.com/nuclio/nuclio-sdk/pkg/logger"
)

// Context is event context
type Context struct {
	// FunctionName is the name of the function
	FunctionName string
	// FunctionVersion is the version of the function
	FunctionVersion string
	// Logger is logger for this context
	Logger logger.Logger
	// EventID is the event ID
	EventID ID
}
