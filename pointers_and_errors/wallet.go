package pointers_and_errors

import "fmt"

// https://pkg.go.dev/fmt#Stringer
type Stringer interface {
	String() string
}

type Bitcoin int

// Creating methods on type declaration
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// When to use pointers:
// rather than taking a copy of the whole Wallet, we instead take a pointer to that wallet
// so that we can change the original values within it

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	fmt.Printf("address of balance from which we will withdraw is %p \n", &w.balance)
	w.balance -= amount
}

// Things to learn to explain:
// - pointers
// - dereference (which is automatic in Go)
