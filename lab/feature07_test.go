package localappb

import "testing"

func TestFeature07(t *testing.T) {
	if Feature07() == "" {
		t.Fatal("empty feature label")
	}
}
