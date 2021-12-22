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

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatalf("did not get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	//t.fatal --> will stop the test if it is called
	//this is because we do not want to make any more asseertions on the error if there isn't one
	//otherwise the test would proceed to the next part and panic because of a nil pointer



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

		assertError(t, err, "Cannot make withdrawal, insufficient funds")
		assertBalance(t, wallet, startingBalance)

	})

}

//updated test to assert on some kind of error message rather than just the existence of an error.
//1.update our helper for a string to compare against
//2.update the caller


//Current error
//got "Woops, you tried to withdraw more than you have", want "Cannot make withdrawal, insufficient funds"