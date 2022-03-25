package http_server_io

import (
	"io"
)

// FileSystemPlayerStore stores score information about players in file
type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	return 0 // TODO
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	// TODO
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(f.database) // TODO: handle error
	return league
}
