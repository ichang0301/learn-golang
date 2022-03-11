// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors

package pointer

import (
	"errors"
	"fmt"
)

var ErrInsufficientFund = errors.New("cannot withdraw, insufficient funds") // The 'var' keyword allows us to define values global to the package.

type Bitcoin int // create new types from existing ones

func (b Bitcoin) String() string { // implement 'Stringer' interface on Bitcoin. Stringer: https://golang.org/pkg/fmt/#Stringer
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount // struct pointer: '(*w).balance += amount' is cumbersome, so go language permits us to write 'w.balance' without an explicit dereference. : https://go.dev/ref/spec#Method_values
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFund
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
