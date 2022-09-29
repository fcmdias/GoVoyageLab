package main

import "math"

type Account struct {
	Amount float64
}

func (a *Account) Deposit(v float64) {
	a.Amount += v
	a.Amount = rounddown(a.Amount)
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

func (a *Account) MultipleYearCal(y int) float64 {

	var res float64
	for i := 1; i < y; i++ {
		v := a.Cal()
		res += v
		a.Deposit(v)
	}
	res += a.Cal()
	return res
}

func rounddown(v float64) float64 {
	return math.Floor(v*100) / 100
}
