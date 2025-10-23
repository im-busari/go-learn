package main

// Why package name is not hello_world
import "fmt"

const spanish = "Spanish"

const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

// The fmt.Println is a side effect (printing to stdout), and the string we send in is our domain.
// What are side effects in Go?
func main() {
	fmt.Println(Hello("world"))
}
