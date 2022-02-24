// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking

// What's the difference between abstraction and generalization: https://stackoverflow.com/questions/19291776/whats-the-difference-between-abstraction-and-generalization
// DRY(Don't Repeat Yourself): https://en.wikipedia.org/wiki/Don't_repeat_yourself
// test double('Test Double' is a generic term for any case where you replace a production object for testing purposes, what are 'dummy', 'fake', 'stubs', 'spies', and 'mocks'): https://martinfowler.com/bliki/TestDouble.html

package mock

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration      time.Duration
	SleepDuration func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepDuration(c.Duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		// }

		// for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprintln(out, finalWord)
}
