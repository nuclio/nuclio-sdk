package nuclio

import "testing"

func TestResponse(t *testing.T) {
	resp := Response{}
	if resp.Body != nil {
		t.Fatal("resp.Body not nil")
	}
}
