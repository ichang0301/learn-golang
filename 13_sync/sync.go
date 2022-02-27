// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/sync

package synchronize

import "sync"

// type Counter struct {
// 	sync.Mutex //  This is bad and wrong.
// 	value      int
// }

type Counter struct {
	mu    sync.Mutex //  Use mutexes for managing state. But when passing ownership of data, use channels. : https://github.com/golang/go/wiki/MutexOrChannel
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	// c.Lock() // Exposing Lock and Unlock is at best confusing but at worst potentially very harmful to your software if callers of your type start calling these methods.
	// defer c.Unlock()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
