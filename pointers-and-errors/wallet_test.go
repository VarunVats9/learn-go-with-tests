package errors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		assertBalance(t, wallet, Bitcoin(10.0))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{amount: Bitcoin(20.0)}
		err := wallet.Withdraw(Bitcoin(10.0))
		assertBalance(t, wallet, Bitcoin(10.0))
		assertNoError(t, err)
	})

	t.Run("withdraw from insufficient fund", func(t *testing.T) {
		wallet := Wallet{amount: Bitcoin(20.0)}
		err := wallet.Withdraw(Bitcoin(30.0))

		assertBalance(t, wallet, Bitcoin(20.0))
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error, but didn't want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get any")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
