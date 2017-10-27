/*
Copyright 2017 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nuclio

import (
	"errors"
	"time"
)

// ErrUnsupported means Event does not support a method
var ErrUnsupported = errors.New("Event does not support this interface")

// SourceInfoProvider is event source provider
type SourceInfoProvider interface {
	// GetClass gets the class of source (sync, async, etc)
	GetClass() string
	// GetKind gets specific kind of source (http, rabbit mq, etc)
	GetKind() string
}

// Event is a nuclio event
type Event interface {
	GetVersion() int
	GetID() ID
	SetID(id ID)
	SetSourceProvider(sourceInfoProvider SourceInfoProvider)
	GetSource() SourceInfoProvider
	GetContentType() string
	GetBody() []byte
	GetSize() int
	GetHeader(key string) interface{}
	GetHeaderByteSlice(key string) []byte
	GetHeaderString(key string) string
	GetHeaders() map[string]interface{}
	GetField(key string) interface{}
	GetFieldByteSlice(key string) []byte
	GetFieldString(key string) string
	GetFieldInt(key string) (int, error)
	GetFields() map[string]interface{}
	GetTimestamp() time.Time
	GetPath() string
	GetURL() string
	GetMethod() string
}
