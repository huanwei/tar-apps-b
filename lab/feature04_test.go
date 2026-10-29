package localappb

import "testing"

func TestFeature04(t *testing.T) {
	if Feature04() == "" {
		t.Fatal("empty feature label")
	}
}
