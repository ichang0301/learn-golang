// https://quii.gitbook.io/learn-go-with-tests/build-an-application/time

package command_line_time

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	store   PlayerStore
	in      *bufio.Scanner
	out     io.Writer
	alerter BlindAlerter
}

// NewCLI creates an CLI
func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		store:   store,
		in:      bufio.NewScanner(in),
		out:     out,
		alerter: alerter,
	}
}

// PlayPoker input '{player name} wins' from user and record win in CLI.store
func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	cli.scheduleBlindAlerts()
	text := cli.readLine()
	cli.store.RecordWin(extractWinner(text)) // `Scanner.Text()` return the string the scanner read to.
}

func (cli *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}

func (cli *CLI) readLine() string {
	cli.in.Scan() // `Scanner.Scan()` will read up to a newline.
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1) // documantation: https://pkg.go.dev/strings#Replace
}
