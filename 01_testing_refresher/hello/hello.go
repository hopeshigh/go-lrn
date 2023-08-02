package hello

func Hello(input string) string {
	if input == "" {
		input = "World"
	}
	return "Hello, " + input + "!"
}
