package command_line

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		player := "Chris"
		AssertPlayerWin(t, playerStore, player)
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		player := "Cleo"
		AssertPlayerWin(t, playerStore, player)
	})
}
