package main

import(
	"errors"
)


var(
	Balance float64
	incorrectInput = errors.New("amount is incorrect")
	incorrectBalance = errors.New("balance is incorrect")
)

func topUpBalance(amount float64) error { 
	if amount <= 0 {
		return incorrectInput
	}
	Balance += amount
	return nil
}

func chargeFromBalance(amount float64) error {
	if amount <= 0 {
		return incorrectInput
	}
	Balance -= amount
	return nil
}

func TopUpAndGetBalance(amount float64) (float64, error) {
    if amount <= 0 {
		return 0, incorrectInput
	}
    if Balance + amount < 0{
        return 0, incorrectBalance
    }
	topUpBalance(amount)
	return Balance, nil
}	

func ChargeFromAndGetBalance(amount float64) (float64, error) {
    if amount <= 0 {
		return 0, incorrectInput
	}
    if Balance - amount < 0{
        return 0, incorrectBalance
    }
	chargeFromBalance(amount)
	return Balance, nil
}