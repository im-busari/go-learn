package main

// Why package name is not hello_world
import "fmt"

const spanish = "Spanish"
const french = "French"

const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

// The fmt.Println is a side effect (printing to stdout), and the string we send in is our domain.
// What are side effects in Go?
func main() {
	fmt.Println(Hello("world", ""))
}
