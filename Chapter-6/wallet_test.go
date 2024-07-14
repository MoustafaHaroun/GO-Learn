package Chapter_6

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		startingBalance := Bitcoin(100)
		wallet := Wallet{balance: startingBalance}

		_ = wallet.Withdraw(Bitcoin(20))
		want := Bitcoin(80)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func TestBitcoin(t *testing.T) {
	t.Run("Stringer_Bitcoin_ShouldReturnString", func(t *testing.T) {
		got := Bitcoin(56).Stringer()
		want := "56 BTC"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Error("wanted an error but didn't get one")
	}

	if !errors.Is(got, want) {
		t.Errorf("got %s want %s", got, want)
	}
}
