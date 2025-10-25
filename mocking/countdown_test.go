package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	// We are using Dependency Injection to facilitate testing the Print to stdout
	// TODO: read more about buffer and bytes
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
