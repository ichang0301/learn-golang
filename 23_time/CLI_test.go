package command_line_time_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	poker "github.com/ichang0301/learn-golang/23_time"
)

var (
	dummyStdIn  = &bytes.Buffer{}
	dummyStdOut = &bytes.Buffer{}
)

type GameSpy struct {
	StartCalled  bool
	StartedWith  int
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.FinishedWith != "Chris" {
			t.Errorf("expected finish called with 'Chris' but got %q", game.FinishedWith)
		}
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.FinishedWith != "Cleo" {
			t.Errorf("expected finish called with 'Cleo' but got %q", game.FinishedWith)
		}
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		player := "Pies"
		stdout := &bytes.Buffer{}
		in := strings.NewReader(fmt.Sprintln(player))
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}

		gotPrompt := stdout.String()
		wantPrompt := fmt.Sprintf("%splease enter the 'number', strconv.Atoi: parsing \"%s\": invalid syntax", poker.PlayerPrompt, player)
		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}
	})
}
