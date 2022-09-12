package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev, _ := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

// to test only this function, run `go test -run=FuzzReverse`
// to run a specific corpus entry with FuzzXxx/testdata eg. go test -run=FuzzReverse/674ae72cbe3ff0c341e8a8fcc254623769b9737d1dff3ded1d7743697cbfbb3f
// The fuzz test will run until it encounters a failing input unless you pass the -fuzztime flag. The default is to run forever if no failures occur, and the process can be interrupted with ctrl-C.
// Fuzz it with go test -fuzz=Fuzz -fuzztime 30s which will fuzz for 30 seconds before exiting if no failure was found
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
