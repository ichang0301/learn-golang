// https://quii.gitbook.io/learn-go-with-tests/questions-and-answers/context-aware-reader

// https://pace.dev/blog/2020/02/03/context-aware-ioreader-for-golang-by-mat-ryer.html

package context_awareness

import (
	"io"
)

func NewCancellableReader(rdr io.Reader) io.Reader {
	return rdr
}
