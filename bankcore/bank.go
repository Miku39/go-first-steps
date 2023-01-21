package bank

import (
	"errors"
	"fmt"
)

// Customer ...
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// Account ...
type Account struct {
	Customer
	Number  int32
	Balance float64
}

// Deposit ...
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

// Withdraw
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw should be greater than account's balance")
	}

	a.Balance -= amount
	return nil
}

// Transfer ...
func (a *Account) Transfer(amount float64, toAccount *Account) error {
	withdrawError := a.Withdraw(amount)
	if withdrawError != nil {
		return withdrawError
	}

	depositError := toAccount.Deposit(amount)
	if depositError != nil {
		deposit2Error := a.Deposit(amount)
		if deposit2Error != nil {
			panic(deposit2Error)
		}
		return depositError
	}

	return nil
}

// Statement ...
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

type Bank interface {
	Statement() string
}

func Statement(b Bank) string {
	return b.Statement()
}
