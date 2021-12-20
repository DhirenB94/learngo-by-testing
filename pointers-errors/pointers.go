package pointers

import "fmt"

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}
//by declaring a new type (Bitcoin) we can now declare methods on it
//we are implementing Stringr on Bitcoin
//Syntax to do this is the same as for a struct
//this interface is defined in fmt, it allows you to define how your type is printed when using %s

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

//Now we use pointers which allows us to actually change the value
//rather than taking a copy of the whole Wallet struct, we instead take a pointer to it so we can change the value
// the difference is the receiver TYPE is now of type *Wallet not type Wallert