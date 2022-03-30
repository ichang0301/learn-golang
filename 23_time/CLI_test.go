package command_line_time_test

import (
	"strings"
	"testing"

	poker "github.com/ichang0301/learn-golang/23_time"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &poker.StubPlayerStore{}

		cli := poker.NewCLI(store, in)
		cli.PlayPoker()

		player := "Chris"
		poker.AssertPlayerWin(t, store, player)
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := &poker.StubPlayerStore{}

		cli := poker.NewCLI(store, in)
		cli.PlayPoker()

		player := "Cleo"
		poker.AssertPlayerWin(t, store, player)
	})
}
