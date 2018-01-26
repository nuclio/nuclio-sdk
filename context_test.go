package nuclio

import (
	"testing"
)

func TestContext(t *testing.T) {
	ctx := Context{}
	if ctx.Logger != nil {
		t.Fatal("Bad Logger")
	}
}
