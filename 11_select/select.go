package selects

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now() // time.Now() record just before we try and get the URL.
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
