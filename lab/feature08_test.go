package localappb

import "testing"

func TestFeature08(t *testing.T) {
	if Feature08() == "" {
		t.Fatal("empty feature")
	}
}
