package http_server_io

import (
	"os"
)

type tape struct {
	// file io.ReadWriteSeeker // io.ReadWriteSeeker is embedding io.Reader, io.Writer and io.Seeker interface
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0) // os.File has a truncate function that will let us effectively empty the file.
	t.file.Seek(0, 0)
	return t.file.Write(p)
}
