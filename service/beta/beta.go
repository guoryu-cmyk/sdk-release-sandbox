package beta

// Hello returns the fixed beta service name used by SDK release smoke tests; callers receive the stable name.
func Hello() string { return "beta" }

// Greeting returns a personalized greeting, or a generic one if no name is given.
func Greeting(name string) string {
	if name == "" {
		return "hello"
	}
	return "hello, " + name
}

// ResourceName returns a service-qualified resource name.
func ResourceName(name string) string {
	if name == "" {
		return Hello()
	}
	return Hello() + "/" + name
}
