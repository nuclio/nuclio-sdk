package nuclio

import "testing"

func TestNewID(t *testing.T) {

	id := NewID()
	if id == nil {
		t.Fatal("id is nil")
	}
}
