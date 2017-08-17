package nuclio

import "testing"

func TestDataBinding(t *testing.T) {
	var db DataBinding

	if db != nil {
		t.Fatal("db not nil")
	}
}
