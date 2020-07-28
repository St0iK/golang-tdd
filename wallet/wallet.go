package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// we can have methods on Types as well !
type Bitcoin int

// Wallet ...
type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

// Deposit when you call a function or a method the arguments are copied.
// Using a pointer here, it will not copy the value, but it will use the existing one
// that will allow us to change the values of the w(Wallet)
// *Wallet -> "Pointer to Wallet"
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	// These pointers to `structs` even have their own name: `struct pointers` and they are automatically dereferenced.
	// this is also valid `(*w).balance`
	w.balance = amount
}

// Balance ...
func (w Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw ...
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}

// Stringer ...
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
