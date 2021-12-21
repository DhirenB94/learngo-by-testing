package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}


	t.Run("Deposit test", func(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin (10))
	assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw test", func(t *testing.T) {
		wallet := Wallet{
			balance: 20,
		}

		wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(15))
	})

	//Returning error when withdrawal exceeds balance
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(100)

		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(200))


		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Errorf("Wanted an error but didnt get one")
		}

		// we want to return an error if you try and withdraw more than you have, and the balance should stay the same
		//we then check an error has returned by failing the test if it is nil
		//Errors can be nil because the return type of Withdraw will be error, which is an interface.
		//functions that take arguments or return values that are interfaces, they can be nillable.

		// Current error
		// ./pointers_test.go:40:25: wallet.Withdraw(Bitcoin(200)) used as value
	})
}
