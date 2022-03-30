// https://quii.gitbook.io/learn-go-with-tests/build-an-application/time

package command_line_time

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
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
	numberOfPlayers, err := strconv.Atoi(cli.readLine())
	if err != nil {
		log.Fatalf("please enter the 'number', %+v", err)
	}

	cli.scheduleBlindAlerts(numberOfPlayers)

	text := cli.readLine()
	cli.store.RecordWin(extractWinner(text)) // `Scanner.Text()` return the string the scanner read to.
}

func (cli *CLI) scheduleBlindAlerts(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (cli *CLI) readLine() string {
	cli.in.Scan() // `Scanner.Scan()` will read up to a newline.
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1) // documantation: https://pkg.go.dev/strings#Replace
}
