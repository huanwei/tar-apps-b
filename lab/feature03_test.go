package localappb

import "testing"

func TestFeature03(t *testing.T) {
	if Feature03() == "" {
		t.Fatal("empty feature label")
	}
}
