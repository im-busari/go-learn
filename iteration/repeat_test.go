package iteration

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a")
	expected := "aaaaa"

	assertCorrectResult(t, result, expected)

}

// Writing benchmarks: https://pkg.go.dev/testing#hdr-Benchmarks
// The testing.B gives you access to the loop function. Loop() returns true as long as the benchmark should continue running.
func BenchmarkRepeat(b *testing.B) {
	// It runs the function nth times to determine the average time it takes to run the function
	for b.Loop() {
		Repeat("a")
	}
}

func assertCorrectResult(t testing.TB, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("expected %q, but got %q", expected, result)
	}
}
