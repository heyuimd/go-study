package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("not enough money")

func NewAccount(owner string) *Account {
	return &Account{owner: owner, balance: 0}
}

func (a *Account) Owner() string {
	return a.owner
}

func (a *Account) Balance() int {
	return a.balance
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil
}

func (a *Account) String() string {
	return fmt.Sprint(a.owner, "'s account has ", a.balance)
}
