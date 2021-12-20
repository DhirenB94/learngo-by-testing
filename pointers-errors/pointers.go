package pointers

import "fmt"

type Wallet struct {
	balance int
}


func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}

//You would think this would return a passing test
//as you are depositing the amount variable (10) to the balance
//then the balance method should return the current state of it

//BUT NO
// In Go, when you call a function or a method, the arguements are copied!
//so when we call func(w Wallet)Deposit(amount int) --> the w is a copy of whatever we called the method from

