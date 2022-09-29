package main

import "testing"

func TestAccount_MultipleYearCal(t *testing.T) {
	type Test struct {
		initialamount float64
		expected      float64
		years         int
	}
	var tests = []Test{
		Test{
			100,
			2.01,
			2,
		},
		Test{
			1000,
			50.6,
			2,
		},
		Test{
			1000,
			20,
			1,
		},
		Test{
			100,
			1,
			1,
		},
		Test{
			3000,
			90,
			1,
		},
		Test{
			3000,
			376.52,
			4,
		},
		Test{
			1000,
			330.82,
			10,
		},
	}

	for _, test := range tests {
		a := Account{test.initialamount}
		result := a.MultipleYearCal(test.years)
		if test.expected != result {
			t.Fatalf("test failed. amount: %v, result: %v, expected: %v", a.Amount, result, test.expected)
		}
	}
}
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
