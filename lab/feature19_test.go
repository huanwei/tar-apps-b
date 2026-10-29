package localappb

import "testing"

func TestFeature19(t *testing.T) {
	if Feature19() == "" {
		t.Fatal("empty feature label")
	}
}
