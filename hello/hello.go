package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello : returns the "hello, world" string
func Hello(name string) string {

	if name == "" {
		name = "world"
	}

	return englishHelloPrefix + name
}

func main() {

	fmt.Println(Hello("world"))
}
