package nuclio

import "testing"

func TestAsync(t *testing.T) {
	var async Async

	if async != nil {
		t.Fatal("async not nil")
	}
}
