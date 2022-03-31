// https://quii.gitbook.io/learn-go-with-tests/build-an-application/websockets

package websockets

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// CLI helps players through a game of poker.
type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

// NewCLI creates a CLI for playing poker.
func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

const (
	PlayerPrompt         = "Please enter the number of players: "                                     // PlayerPrompt is the text asking the user for the number of players.
	BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number" // BadPlayerInputErrMsg is the text telling the user they did bad things.
	BadWinnerInputMsg    = "invalid winner input, expect format of 'PlayerName wins'"                 // BadWinnerInputMsg is the text telling the user they declared the winner wrong.
)

// PlayPoker starts to play the game.
func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayers, err := strconv.Atoi(cli.readLine())
	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers, cli.out)

	winnerInput := cli.readLine()
	winner, err := extractWinner(winnerInput)
	if err != nil {
		fmt.Fprint(cli.out, BadWinnerInputMsg)
		return
	}

	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", errors.New(BadWinnerInputMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}
