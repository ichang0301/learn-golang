package command_line_time_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/ichang0301/learn-golang/23_time"
)

var (
	dummyBlindAlerter = &SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	dummyStdIn        = &bytes.Buffer{}
	dummyStdOut       = &bytes.Buffer{}
)

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		store := &poker.StubPlayerStore{}
		game := poker.NewGame(dummyBlindAlerter, store)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		player := "Chris"
		poker.AssertPlayerWin(t, store, player)
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")
		store := &poker.StubPlayerStore{}
		game := poker.NewGame(dummyBlindAlerter, store)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		player := "Cleo"
		poker.AssertPlayerWin(t, store, player)
	})
}

func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	t.Helper()

	if got.amount != want.amount {
		t.Errorf("got amount %d, want %d", got.amount, want.amount)
	}

	if got.at != want.at {
		t.Errorf("got scheduled time of %v, want %v", got.at, want.at)
	}
}
