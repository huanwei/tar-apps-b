package localappb

import "testing"

func TestFeature14(t *testing.T) {
	if Feature14() == "" {
		t.Fatal("empty feature label")
	}
}
