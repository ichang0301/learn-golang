package http_server_io

import (
	"encoding/json"
	"io"
	"log"
)

// FileSystemPlayerStore stores score information about players in file
type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker // io.ReadSeeker is embedding io.Reader and io.Seeker interface
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	f.Database.Seek(0, 0)
	json.NewEncoder(f.Database).Encode(&league)
}

func (f *FileSystemPlayerStore) GetLeague() League {
	if _, err := f.Database.Seek(0, io.SeekStart); err != nil { // go back to the start.
		log.Fatal(err)
	}
	league, _ := NewLeague(f.Database) // TODO: handle error
	return league
}
