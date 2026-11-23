package alpha

import "testing"

func TestHello(t *testing.T) {
	if got := Hello(); got != "alpha" {
		t.Fatalf("Hello() = %q, want alpha", got)
	}
}
