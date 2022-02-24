package selects

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select { // What select lets you do is wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} { // A 'chan struct{}' is the smallest data type available from a memory perspective
	ch := make(chan struct{})
	// var ch chan struct{}	// [Caution] This occurs error 'panic: close of nil channel'. When you use 'var', the variable will be initialized with 'nil'. : https://go.dev/play/p/IIbeAox5jKA
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now() // time.Now() record just before we try and get the URL.
// 	http.Get(url)
// 	return time.Since(start)
// }
