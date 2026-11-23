package beta

import "testing"

func TestHello(t *testing.T) {
	if got := Hello(); got != "beta" {
		t.Fatalf("Hello() = %q, want beta", got)
	}
}
