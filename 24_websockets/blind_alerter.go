package websockets

import (
	"fmt"
	"io"
	"time"
)

// BlindAlerter schedules alerts for blind amounts.
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

// BlindAlerterFunc allows you to implement BlindAlerter with a function. (For example after implement StdOutAlerter, and then you can use type conversion like `BlindAlerterFunc(StdOutAlerter)`)
type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

// ScheduleAlertAt is BlindAlerterFunc implementation of BlindAlerter.
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) { // If you are making a library that exposes an interface with one function defined it is a common idiom to also expose a MyInterfaceFunc type, so rather than having to create an empty struct type.
	a(duration, amount, to)
}

// Alerter will schedule alerts and print them to os.Stdout.
func Alerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() { // `time.After(duration)` returns a `chan Time` when the duration has expired. documantation: https://pkg.go.dev/time#AfterFunc, `time.NewTicker(duration)` returns a 'Ticker' channel every duration, rather than just once.
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}
