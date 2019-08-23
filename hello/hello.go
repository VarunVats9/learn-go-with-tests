package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello : returns the "hello, world" string
func Hello(name, language string) string {

	if name == "" {
		name = "world"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name
	}

	if language == "French" {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {

	fmt.Println(Hello("world", ""))
}
