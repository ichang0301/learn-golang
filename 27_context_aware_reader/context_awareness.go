// https://quii.gitbook.io/learn-go-with-tests/questions-and-answers/context-aware-reader

// https://pace.dev/blog/2020/02/03/context-aware-ioreader-for-golang-by-mat-ryer.html

package context_awareness

import (
	"context"
	"io"
)

type readerCtx struct { // delegation pattern: https://en.wikipedia.org/wiki/Delegation_pattern
	ctx      context.Context
	delegate io.Reader
}

func (r *readerCtx) Read(p []byte) (n int, err error) { // implement io.Reader
	// select {
	// case <-r.ctx.Done():	// This is not good, because it only checks ctx is cancelled.
	// 	return 0, r.ctx.Err()
	// default:
	// }
	if err := r.ctx.Err(); err != nil { // return the error from the context.Context, and this allows callers of the code to inspect the various reasons cancellation has occurred and this is covered more in the original post.
		return 0, err
	}
	return r.delegate.Read(p)
}

func NewCancellableReader(ctx context.Context, rdr io.Reader) io.Reader {
	return &readerCtx{
		ctx:      ctx,
		delegate: rdr,
	}
}
