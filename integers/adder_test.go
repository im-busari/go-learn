package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	assertCorrectValue(t, sum, expected)
}

// Examples without output comments are useful for demonstrating code that cannot run as unit tests,
// such as that which accesses the network, while guaranteeing the example at least compiles.
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectValue(t testing.TB, sum, expected int) {
	t.Helper()

	if sum != expected {
		// %d when you want to print integers
		// %q when you want to print strings
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
