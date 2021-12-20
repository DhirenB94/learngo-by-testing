package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	fmt.Printf("address of balance in test is %v\n", &wallet.balance)
	want := 10
//the addresses of the balance field is different here in the test (accessing from Wallet struct
//compared to the balance field used in the function (accessing from w Wallet struct)
//this is because it uses a copy n the function


	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
