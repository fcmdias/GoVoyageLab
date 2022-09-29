package main

import "testing"

func TestAccount_Cal(t *testing.T) {
	type Test struct {
		amount             float64
		expected           float64
		secondyearexpected float64
	}
	var tests = []Test{
		Test{
			100,
			1,
			1.01,
		},
		Test{
			1000,
			20,
			30.60,
		},
		Test{
			3000,
			90,
			92.7,
		},
	}

	for _, test := range tests {
		a := Account{}
		a.Deposit(test.amount)

		if a.Cal() != test.expected {
			t.Fatalf("test failed. amount: %v, result: %v, expected: %v", a.Amount, a.Cal(), test.expected)
		}

		a.Deposit(a.Cal())
		if a.Cal() != test.secondyearexpected {
			t.Fatalf("test failed. amount: %v, result: %v, expected: %v", a.Amount, a.Cal(), test.secondyearexpected)
		}

	}
}
