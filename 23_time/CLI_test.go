package command_line_time_test

import (
	"strings"
	"testing"
	"time"

	poker "github.com/ichang0301/learn-golang/23_time"
)

type SpyBlindAlerter struct {
	alerts []struct {
		scheduledAt time.Duration
		amount      int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		scheduledAt time.Duration
		amount      int
	}{duration, amount})
}

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &poker.StubPlayerStore{}
		dummySpyAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(store, in, dummySpyAlerter)
		cli.PlayPoker()

		player := "Chris"
		poker.AssertPlayerWin(t, store, player)
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := &poker.StubPlayerStore{}
		dummySpyAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(store, in, dummySpyAlerter)
		cli.PlayPoker()

		player := "Cleo"
		poker.AssertPlayerWin(t, store, player)
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) != 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})
}
