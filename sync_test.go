package nuclio

import "testing"

func TestSync(t *testing.T) {
	var sync Sync

	if sync != nil {
		t.Fatal("sync not nil")
	}
}

func TestAbstractSync(t *testing.T) {
	as := &AbstractSync{}
	var sync Sync = as // Make sure AbstractSync implements Sync
	if sync.(*AbstractSync) == nil {
		t.Fatal("Bad conversion")
	}
}
