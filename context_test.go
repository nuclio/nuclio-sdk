package nuclio

import "testing"

func TestContext(t *testing.T) {
	var ctx Context

	if ctx.Logger != nil {
		t.Fatal("ctx.Logger not nil")
	}
}
