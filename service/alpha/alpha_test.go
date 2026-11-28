package alpha

import "testing"

func TestHello(t *testing.T) {
	if got := Hello(); got != "alpha" {
		t.Fatalf("Hello() = %q, want alpha", got)
	}
}

func TestResourceName(t *testing.T) {
	for _, tc := range []struct {
		name string
		want string
	}{
		{name: "", want: "alpha"},
		{name: "demo", want: "alpha/demo"},
        {name: "hwtest", want: "alpha/hwtest"},
        {name: "hwtest1", want: "alpha/hwtest1"},
	} {
		if got := ResourceName(tc.name); got != tc.want {
			t.Errorf("ResourceName(%q) = %q, want %q", tc.name, got, tc.want)
		}
	}
}
