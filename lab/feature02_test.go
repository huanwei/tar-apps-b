package localappb

import "testing"

func TestFeature02(t *testing.T) {
	if Feature02() == "" {
		t.Fatal("empty feature")
	}
}
