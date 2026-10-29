package localappb

import "testing"

func TestFeature16(t *testing.T) {
	if Feature16() == "" {
		t.Fatal("empty feature-16")
	}
}
