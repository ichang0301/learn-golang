package http_server_io

import (
	"encoding/json"
	"os"
)

// FileSystemPlayerStore stores score information about players in file
type FileSystemPlayerStore struct {
	Database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, 0) // go back to the start.
	league, _ := NewLeague(file)
	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&tape{file}),
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

	f.Database.Encode(&f.league)
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}
