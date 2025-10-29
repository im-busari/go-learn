package _select

import (
	"net/http"
	"time"
)

// You have been asked to make a function called WebsiteRacer which takes 2 URLs and "races"
// them by hitting them within 10 seconds then it should return an error
func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)

	return time.Since(start)
}
