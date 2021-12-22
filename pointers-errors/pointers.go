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

func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
		return errors.New("Woops, you tried to withdraw more than you have")
	}
	w.balance -= amount
	return nil
}