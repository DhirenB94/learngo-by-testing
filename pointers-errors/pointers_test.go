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


		assertBalance(t, wallet, startingBalance) //want = startingBalance, because you dont want anything to actually be withdrawn if its too much

		if err == nil {
			t.Errorf("Wanted an error but didnt get one")
		}

		// we want to return an 'error' if you try and withdraw more than you have, and the balance should stay the same

		//current error
		// Tests runs but fails
		//		pointers_test.go:43: got -100 whatever the hell you want, want 100 whatever the hell you want --> it is withdrawing
		//    	pointers_test.go:46: Wanted an error but didnt get one --> because we've told withdraw func returns nil which is triggering the err

	})
}
