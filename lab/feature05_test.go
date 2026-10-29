package localappb

import "testing"

func TestFeature05(t *testing.T) {
	if Feature05() == "" {
		t.Fatal("empty feature")
	}
}
