package websockets

import (
	"os"
)

type Tape struct {
	// File io.ReadWriteSeeker // io.ReadWriteSeeker is embedding io.Reader, io.Writer and io.Seeker interface
	File *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.File.Truncate(0) // os.File has a truncate function that will let us effectively empty the file.
	t.File.Seek(0, 0)
	return t.File.Write(p)
}
