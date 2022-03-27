package http_server_io

import (
	"io"
	"log"
)

// FileSystemPlayerStore stores score information about players in file
type FileSystemPlayerStore struct {
	database io.ReadSeeker // io.ReadSeeker is embedding io.Reader and io.Seeker interface
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	// TODO
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	if _, err := f.database.Seek(0, io.SeekStart); err != nil { // go back to the start.
		log.Fatal(err)
	}
	league, _ := NewLeague(f.database) // TODO: handle error
	return league
}
