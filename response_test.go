package nuclio

import (
	"testing"
)

func TestResponse(t *testing.T) {
	resp := Response{}
	if resp.StatusCode != 0 {
		t.Fatal("Bad StatusCode")
	}
}
