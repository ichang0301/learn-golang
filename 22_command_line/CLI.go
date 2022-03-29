package command_line

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	in    io.Reader
}

func (cli *CLI) PlayPoker() {
	scanner := bufio.NewScanner(cli.in) // documantation: https://pkg.go.dev/bufio#Scanner
	scanner.Scan()                      // `Scanner.Scan()` will read up to a newline.

	cli.store.RecordWin(extractWinner(scanner.Text())) // `Scanner.Text()` return the string the scanner read to.
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1) // documantation: https://pkg.go.dev/strings#Replace
}
