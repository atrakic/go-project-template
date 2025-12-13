package greeting

func Hello() string {
	return "Hello, world."
}

func HelloWithName(name string) string {
	if name == "" {
		return Hello()
	}
	return "Hello, " + name + "."
}
