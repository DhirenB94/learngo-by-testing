package pointers

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