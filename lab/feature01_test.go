package localappb

import "testing"

func TestFeature01(t *testing.T) {
	if Feature01() == "" {
		t.Fatal("empty feature")
	}
}
