package main

// Why package name is not hello_world
import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

// The fmt.Println is a side effect (printing to stdout), and the string we send in is our domain.
// What are side effects in Go?
func main() {
	fmt.Println(Hello("world"))
}
