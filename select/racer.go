package _select

import (
	"net/http"
	"time"
)

// You have been asked to make a function called WebsiteRacer which takes 2 URLs and "races"
// them by hitting them within 10 seconds then it should return an error
func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)

	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}

	return b
}
