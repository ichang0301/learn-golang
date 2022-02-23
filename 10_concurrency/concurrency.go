// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency

// Make it work, make it right, make it fast -  Kent Beck: http://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast
// Premature optimization is the root of all evil -- Donald Knuth: http://wiki.c2.com/?PrematureOptimization

package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string // As we don't need either value to be named, each of them is anonymous within the struct
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// The problem here is that the variable url is reused for each iteration of the for loop - it takes a new value from urls each time. Let's study about 'lexical(=static) scope vs dynamic scope' again. we can find this using `go vet`.
	// for _, url := range urls {
	// go func() {
	// 	results[url] = wc(url)
	// }()
	// }

	// This is a race condition, a bug that occurs when the output of our software is dependent on the timing and sequence of events that we have no control over. Go can help us to spot race conditions with its built in race detector(`go test -race`): https://go.dev/blog/race-detector
	// for _, url := range urls {
	// go func(u string) {
	// 	results[u] = wc(u)
	// }(url)
	// }

	// We can solve this data race by coordinating our goroutines using channels.
	for _, url := range urls {
		go func(u string) { // anonymous function
			resultChannel <- result{u, wc(u)} // 'send statement'
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel // 'receive statement'
		results[r.string] = r.bool
	}

	return results
}
