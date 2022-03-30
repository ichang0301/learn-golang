// https://quii.gitbook.io/learn-go-with-tests/build-an-application/time

package command_line_time

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

// CLI helps players through a game of poker.
type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game *TexasHoldem
}

// NewCLI creates a CLI for playing poker.
func NewCLI(in io.Reader, out io.Writer, game *TexasHoldem) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

// PlayerPrompt is the text asking the user for the number of players.
const PlayerPrompt = "Please enter the number of players: "

// PlayPoker starts to play the game.
func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayers, err := strconv.Atoi(cli.readLine())
	if err != nil {
		log.Fatalf("please enter the 'number', %+v", err)
	}

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)
	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
