package localappb

import "testing"

func TestFeature06(t *testing.T) {
	if Feature06() == "" {
		t.Fatal("empty feature")
	}
}
