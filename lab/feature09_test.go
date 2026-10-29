package localappb

import "testing"

func TestFeature09(t *testing.T) {
	if Feature09() == "" {
		t.Fatal("empty feature-09")
	}
}
