// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency

// Make it work, make it right, make it fast -  Kent Beck: http://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast
// Premature optimization is the root of all evil -- Donald Knuth: http://wiki.c2.com/?PrematureOptimization

package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
