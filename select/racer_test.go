package _select

import "testing"

func TestRacer(t *testing.T) {
	slowURL := ""
	fastURL := ""

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
