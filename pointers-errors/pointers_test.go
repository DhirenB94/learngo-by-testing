package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit test", func(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin (10))
	assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw test", func(t *testing.T) {
		wallet := Wallet{
			balance: 20,
		}

		err := wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(15))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(100)

		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(200))

		assertError(t, err, insufficientFundsError)
		assertBalance(t, wallet, startingBalance)

	})
}
func assertBalance (t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError (t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("did not get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
// Check if withdraw is successful, an error is NOT returned
func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatalf("did not get an error but wanted one")
	}
}
