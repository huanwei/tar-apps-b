package localappb

import "testing"

func TestFeature20(t *testing.T) {
	if Feature20() == "" {
		t.Fatal("empty feature label")
	}
}
