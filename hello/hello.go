package main

import "fmt"

// Hello : returns the "hello, world" string
func Hello(name string) string {

	return "Hello, " + name
}

func main() {

	fmt.Println(Hello("world"))
}
