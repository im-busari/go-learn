package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// An operation that does not block in Go will run in a separate process called a goroutine

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// Channels are a Go data structure that can both receive and send values. These operations,
	// along with their details, allow communication between different processes.
	resultChannel := make(chan result)

	for _, url := range urls {
		// Because the only way to start a goroutine is to put go in front of a function call,
		// we often use anonymous functions when we want to start a goroutine
		go func() {
			//  results[url] = wc(url)

			// Send statement (<-)
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression
		r := <-resultChannel

		results[r.string] = r.bool
	}

	//time.Sleep(2 * time.Second)

	return results
}
