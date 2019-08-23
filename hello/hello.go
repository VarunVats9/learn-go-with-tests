package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello : returns the "hello, world" string
func Hello(name, language string) string {

	if name == "" {
		name = "world"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}

	return englishHelloPrefix + name
}

func main() {

	fmt.Println(Hello("world", ""))
}
