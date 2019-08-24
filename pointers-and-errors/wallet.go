package errors

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds is the error message
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient amount")

// Bitcoin of type float64
type Bitcoin float64

// Wallet is the struct which stores the bitcoin field.
type Wallet struct {
	amount Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

// Deposit takes the bitcoin field for the wallet receiver.
func (w *Wallet) Deposit(deposit Bitcoin) {
	w.amount += deposit
}

// Balance returns the bitcoin field for the wallet receiver.
func (w *Wallet) Balance() Bitcoin {
	return w.amount
}

// Withdraw removes the given bitcoin field for the wallet receiver.
func (w *Wallet) Withdraw(withdraw Bitcoin) error {
	if w.amount < withdraw {
		return ErrInsufficientFunds
	}
	w.amount -= withdraw
	return nil
}
