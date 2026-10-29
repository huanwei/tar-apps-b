package localappb

import "testing"

func TestFeature10(t *testing.T) {
	if Feature10() == "" {
		t.Fatal("empty feature")
	}
}
