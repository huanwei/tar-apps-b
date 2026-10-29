package localappb

import "testing"

func TestFeature17(t *testing.T) {
	if Feature17() == "" {
		t.Fatal("empty feature")
	}
}
