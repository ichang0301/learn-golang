package selects

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select { // What select lets you do is wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout): // 'time.After' returns a chan (like ping) and will send a signal down it after the amount of time you define. We can use 'time.After' to prevent your system blocking forever.
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
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
