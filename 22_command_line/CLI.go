package command_line

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	in    *bufio.Scanner // documantation: https://pkg.go.dev/bufio#Scanner
}

// NewCLI creates an CLI
func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		store: store,
		in:    bufio.NewScanner(in),
	}
}

// PlayPoker input '{player name} wins' from user and record win in CLI.store
func (cli *CLI) PlayPoker() {
	text := cli.readLine()
	cli.store.RecordWin(extractWinner(text)) // `Scanner.Text()` return the string the scanner read to.
}

func (cli *CLI) readLine() string {
	cli.in.Scan() // `Scanner.Scan()` will read up to a newline.
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1) // documantation: https://pkg.go.dev/strings#Replace
}
