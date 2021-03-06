package http_server_io

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// FileSystemPlayerStore stores score information about players in file
type FileSystemPlayerStore struct {
	Database *json.Encoder // When you rely on global variables responsibilities become very unclear.
	league   League
}

// NewFileSystemPlayerStore creates an FileSystemPlayerStore that implements the PlayerStore interface
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	if err := initialisePlayerDBFile(file); err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0) // go back to the start.

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
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
	sort.Slice(f.league, func(i, j int) bool { // documantation: https://pkg.go.dev/sort#Slice
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}
