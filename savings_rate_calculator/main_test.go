package main

import "testing"

func TestAccount_Cal(t *testing.T) {
	type Test struct {
		amount   float64
		expected float64
	}
	var tests = []Test{
		Test{
			100,
			1,
		},
		Test{
			1000,
			20,
		},
		Test{
			3000,
			90,
		},
	}

	for _, test := range tests {
		a := Account{test.amount}

		if a.Cal() != test.expected {
			t.Fatalf("test failed. amount: %v, result: %v, expected: %v", a.Amount, a.Cal(), test.expected)
		}

	}
}
