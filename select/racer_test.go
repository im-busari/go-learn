package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the URL of the dfastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// By prefixing a function call with defer it will now call that function at the end
		// of the containing function.
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL // "http://facebook.com"
		fastURL := fastServer.URL // "http://www.quii.dev"

		want := fastURL
		got, _ := RacerWithSelect(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(25 * time.Millisecond)
		serverB := makeDelayedServer(25 * time.Millisecond)

		defer serverB.Close()
		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an erro but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
