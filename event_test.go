package nuclio

import "testing"

func TestEvent(t *testing.T) {
	var si SourceInfoProvider

	if si != nil {
		t.Fatal("si not nil")
	}

	var evt Event
	if evt != nil {
		t.Fatal("evt not nil")
	}
}

func TestAbstractEvent(t *testing.T) {
	ae := &AbstractEvent{}

	var evt Event = ae // Make sure AbstractEvent implement Event interface
	if evt.(*AbstractEvent) == nil {
		t.Fatal("nil event")
	}
}
