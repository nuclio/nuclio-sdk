package nuclio

import (
	"errors"
	"net/http"
	"testing"
)

func TestErrors(t *testing.T) {
	errorMessage := "Object created"

	// Make sure it implements WithStatusCode interface
	errCreated := ErrCreated
	errCreated.error = errors.New(errorMessage)
	var withStatus WithStatusCode = &ErrCreated

	if withStatus.StatusCode() != http.StatusCreated {
		t.Fatal("Bad status")
	}

	// Make sure it implements Error interface
	var err error = &errCreated
	if err.Error() != errorMessage {
		t.Fatal("Bad error message")
	}
}
