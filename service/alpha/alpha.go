package alpha

// Hello returns the fixed alpha service name used by SDK release smoke tests; callers receive the stable name.
func Hello() string { return "alpha" }

// ResourceName returns a service-qualified resource name.
func ResourceName(name string) string {
	if name == "" {
		return Hello()
	}
	return Hello() + "/" + name
}
