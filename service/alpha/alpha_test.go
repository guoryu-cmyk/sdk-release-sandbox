package alpha

import "testing"

func TestHello(t *testing.T) {
	if Hello() != "alpha" {
		t.Fatal("unexpected")
	}
}
