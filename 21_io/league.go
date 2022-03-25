package http_server_io

import (
	"encoding/json"
	"fmt"
	"io"
)

// NewLeague read JSON data and return league information and error
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}
