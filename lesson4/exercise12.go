package main

import (
	"errors"
	"lesson/bankUsers"
)

type Account struct {
	balance float64
	owner   string
}

func NewAccount(owner string) *Account {
	return &Account{balance: bankUsers.NewConfig().DefaultBalance, owner: owner}
}

func (a *Account) SetBalance(value float64) error {
	if value < 0 {
		return errors.New("Ошибка при установке баланса")
	}

	a.balance = value
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(value float64) error {
	if value <= 0 {
		return errors.New("Ошибка при внесении денег")
	}

	a.balance += value
	return nil
}

func (a *Account) Withdraw(value float64) error {
	if value <= 0 {
		return errors.New("Ошибка при снятии денег")
	} else if value > a.balance {
		return errors.New("Ошибка при снятии денег")
	}

	a.balance -= value
	return nil
}
