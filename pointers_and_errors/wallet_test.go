package pointers_and_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(10)

	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	if got != want {
		t.Errorf("Expected %d, but got %d", want, got)
	}

	if got != want {
		// This will use the String() function on the type (10 BTC)
		t.Errorf("Expected %s, but got %s", want, got)
	}
}
