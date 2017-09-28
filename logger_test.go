package nuclio

import "testing"

func TestLogger(t *testing.T) {
	var logger Logger

	if logger != nil {
		t.Fatalf("Non nil logger")
	}
}
