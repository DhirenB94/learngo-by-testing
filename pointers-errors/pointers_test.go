package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit test", func(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin (10))

	got := wallet.Balance()
	want := Bitcoin(20)


	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	})

	t.Run("Withdraw test", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Withdraw(Bitcoin(5))

		got:= wallet.Balance()
		want := Bitcoin(5)

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
