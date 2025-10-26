package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	// We are using Dependency Injection to facilitate testing the Print to stdout
	// TODO: read more about buffer and bytes
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}
