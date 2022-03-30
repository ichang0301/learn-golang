package command_line_time

import "time"

// Game manages the state of a game.
type Game struct {
	store   PlayerStore
	alerter BlindAlerter
}

// NewGame creates a Game for managing poker.
func NewGame(alerter BlindAlerter, store PlayerStore) *Game {
	return &Game{
		alerter: alerter,
		store:   store,
	}
}

// Start starts poker with the number of players.
func (g *Game) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

// Finish records winner.
func (p *Game) Finish(winner string) {
	p.store.RecordWin(winner)
}
