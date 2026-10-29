package localappb

import "testing"

func TestFeature18(t *testing.T) {
	if Feature18() == "" {
		t.Fatal("empty feature label")
	}
}
