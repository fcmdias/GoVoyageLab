package main

import "math"

type Account struct {
	Amount float64
}

func (a *Account) Deposit(v float64) {
	a.Amount += v
}

func (a Account) Cal() float64 {

	var result float64 = a.Amount * 0.01
	if a.Amount > 1000 {
		result = a.Amount * 0.03
	} else if a.Amount > 500 {
		result = a.Amount * 0.02
	} else {

	}

	return math.Floor(result*100) / 100
}
