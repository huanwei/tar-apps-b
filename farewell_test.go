package localappb

import "testing"

func TestFarewell(t *testing.T) {
	if Farewell() == "" {
		t.Fatal("empty farewell")
	}
}
