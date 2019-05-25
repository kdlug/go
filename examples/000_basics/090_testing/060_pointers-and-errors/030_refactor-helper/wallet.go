// go get -u github.com/kisielk/errcheck
// errcheck .
// bashrc
// export GOPATH=/home/user/go
// export PATH=$PATH:$GOPATH/bin

package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFund = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		// errors.New creates a new error with a message
		// return errors.New("cannot withdraw, insufficient funds")
		return ErrInsufficientFund
	}

	w.balance -= amount

	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BCT", b)
}
