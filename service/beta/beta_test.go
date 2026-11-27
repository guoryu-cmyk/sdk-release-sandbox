package beta

import "testing"

func TestHello(t *testing.T) {
	if got := Hello(); got != "beta" {
		t.Fatalf("Hello() = %q, want beta", got)
	}
}

func TestGreeting(t *testing.T) {
	for _, tc := range []struct {
		name string
		want string
	}{
		{name: "", want: "hello"},
		{name: "SDK", want: "hello, SDK"},
	} {
		if got := Greeting(tc.name); got != tc.want {
			t.Errorf("Greeting(%q) = %q, want %q", tc.name, got, tc.want)
		}
	}
}

func TestResourceName(t *testing.T) {
	for _, tc := range []struct {
		name string
		want string
	}{
		{name: "", want: "beta"},
		{name: "demo", want: "beta/demo"},
        {name: "service", want: "beta/service"},
	} {
		if got := ResourceName(tc.name); got != tc.want {
			t.Errorf("ResourceName(%q) = %q, want %q", tc.name, got, tc.want)
		}
	}
}
