package pointers

import "fmt"

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d whatever the hell you want", b)
}
//by declaring a new type (Bitcoin) we can now declare methods on it
//we are implementing Stringr on Bitcoin
//Syntax to do this is the same as for a struct
//this interface is defined in fmt, it allows you to define how your type is printed when used with the %s format strings in print

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

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}