package websockets

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

// Find find and return player in league, if player don't exist in league return nil
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

// NewLeague read JSON data and return league information and error
func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}
