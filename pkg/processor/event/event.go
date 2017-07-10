package event

import (
	"time"
)

// SourceInfoProvider gives information about source provider
type SourceInfoProvider interface {
	// GetClass returns the class of source (sync, async, etc)
	GetClass() string

	// GetKins returns the specific kind of source (http, rabbit mq, etc)
	GetKind() string
}

// Event is a nuclio Event
type Event interface {
	// GetVersion return the event version
	GetVersion() int
	// GetID returns the event ID
	GetID() ID
	// SetID sets the event ID
	SetID(id ID)
	// SetSourceProvider sets the source provider
	SetSourceProvider(sourceInfoProvider SourceInfoProvider)
	// GetSource return the source information
	GetSource() SourceInfoProvider
	// GetContentType returns the event content type
	GetContentType() string
	// GetBody returns the event body
	GetBody() []byte
	// GetSize return the body size in bytes
	GetSize() int
	// GetHeader returns a specific header, nil if not found
	GetHeader(key string) interface{}
	// GetHeaderByteSlice returns header as []byte
	GetHeaderByteSlice(key string) []byte
	// GetHeaderString returns header as string
	GetHeaderString(key string) string
	// GetHeaders returns all headers
	GetHeaders() map[string]interface{}
	// GetTimestamp return the time the event was received
	GetTimestamp() time.Time
	// GetPath returns the path
	GetPath() string
	// GetURL return the URL
	GetURL() string
}
