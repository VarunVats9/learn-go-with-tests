package errors

import "fmt"

// Bitcoin of type float64
type Bitcoin float64

// Wallet is the struct which stores the bitcoin field.
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

// Deposit takes the bitcoin field for the wallet receiver.
func (w *Wallet) Deposit(balance Bitcoin) {
	w.balance += balance
}

// Balance returns the bitcoin field for the wallet receiver.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
