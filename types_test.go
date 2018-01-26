package nuclio

import (
	"testing"
)

func TestID(t *testing.T) {
	id := NewID()
	if len(id.String()) == 0 {
		t.Fatal("empty id")
	}
}
