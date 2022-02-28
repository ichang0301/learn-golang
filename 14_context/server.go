// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/context

// go doc: https://pkg.go.dev/context
// go blog: https://go.dev/blog/context

// 'Context.Value' should inform, not control. If a function needs some values, put them as typed parameters rather than trying to fetch them from context.Value.

package contexts

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context()) // The server code has become simplified as it's no longer explicitly responsible for cancellation, it simply passes through context and relies on the downstream functions to respect any cancellations that may occur.
		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}
