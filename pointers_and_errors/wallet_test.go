package pointers_and_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		if got != want {
			t.Errorf("Expected %d, but got %d", want, got)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(25)}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(15)

		if got != want {
			// This will use the String() function on the type (10 BTC)
			t.Errorf("Expected %s, but got %s", want, got)
		}
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("Wanted an error but didn't get one")
		}
	})
}
