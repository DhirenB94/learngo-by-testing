package pointers

import "fmt"

type Wallet struct {
	balance int
}


func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}

//Now we use pointers which allows us to actually change the value
//rather than taking a copy of the whole Wallet struct, we instead take a pointer to it so we can change the value
// the difference is the receiver TYPE is now of type *Wallet not type Wallert