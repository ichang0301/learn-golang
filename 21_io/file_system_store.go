package http_server_io

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore stores score information about players in file
type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker // io.ReadSeeker is embedding io.Reader, io.Writer and io.Seeker interface
	league   League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0) // go back to the start.
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		Database: database,
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	json.NewEncoder(f.Database).Encode(&f.league)
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}
