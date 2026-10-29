package localappb

import "testing"

func TestFeature11(t *testing.T) {
	if Feature11() == "" {
		t.Fatal("empty feature label")
	}
}
