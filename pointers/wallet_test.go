package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

    err := wallet.Withdraw(Bitcoin(10))

    assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
    startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}

    err := wallet.Withdraw(Bitcoin(100))

    assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()
		if got != want {
			t.Errorf("want %s but got %s\n", want, got)
		}
	}

func assertError(t *testing.T, got, want error) {
    t.Helper()
    if got  == nil {
			t.Fatal("wanted an error but didn't get one")
    }

    if got != want {
			t.Errorf("wanted %q, but got %q", want, got)
    }
  }

