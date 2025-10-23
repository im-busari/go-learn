package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to Chris", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in Spanish", func(t *testing.T) {
		got := Hello("Elly", "Spanish")
		want := "Holla, Elly"

		assertCorrectMessage(t, got, want)
	})
}

// For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B
// both satisfy, so you can call helper functions from a test, or a benchmark
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // needed to tell the test suite that this method is a helper

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
