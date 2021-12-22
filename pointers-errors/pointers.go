package pointers

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d whatever the hell you want", b)
}


type Bitcoin int

type Wallet struct {
	balance Bitcoin
}


func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var insufficientFundsError = errors.New("cannot make withdrawal, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
		return insufficientFundsError
	}
	w.balance -= amount
	return nil
}