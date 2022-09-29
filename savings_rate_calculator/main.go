package main

type Account struct {
	Amount float64
}

func (a *Account) Deposit(v float64) {
	a.Amount += v
}

func (a Account) Cal() float64 {

	if a.Amount > 1000 {
		return a.Amount * 0.03
	} else if a.Amount > 500 {
		return a.Amount * 0.02
	}

	return a.Amount * 0.01
}
