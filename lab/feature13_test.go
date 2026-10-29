package localappb

import "testing"

func TestFeature13(t *testing.T) {
	if Feature13() == "" {
		t.Fatal("empty feature")
	}
}
