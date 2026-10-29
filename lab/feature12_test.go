package localappb

import "testing"

func TestFeature12(t *testing.T) {
	if Feature12() == "" {
		t.Fatal("empty feature label")
	}
}
