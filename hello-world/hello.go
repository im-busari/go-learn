package main

// Why package name is not hello_world
import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

// The fmt.Println is a side effect (printing to stdout), and the string we send in is our domain.
// What are side effects in Go?
func main() {
	fmt.Println(Hello("world"))
}
