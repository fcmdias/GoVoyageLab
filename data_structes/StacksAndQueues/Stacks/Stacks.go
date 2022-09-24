package main

import (
	"errors"
	"fmt"
)

// ErrEmptyStack when Stack is empty
var ErrEmptyStack = errors.New("Empty Stack")

// Stack represents stack that hold a slice
// Last in first out (like a stack of books)
type Stack struct {
	items []int
}

// Push will add a value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end
// and RETURNs the removed value
func (s *Stack) Pop() (int, error) {
	l := len(s.items)
	if l > 0 {
		toRemove := s.items[l-1]
		s.items = s.items[:l-1]
		return toRemove, nil
	}
	return 0, fmt.Errorf("cannot remove value: %w", ErrEmptyStack)
	// return 0, errors.New("Stack is empty therefore cannot remove value using the pop method")
}

func main() {
	s := Stack{}
	s.Push(7)
	s.Push(9)
	s.Push(4)
	s.Push(8)
	fmt.Println(s)
	v, err := s.Pop()
	if err != nil {
		// ignores if EmptyStackErr exists in err
		if !errors.Is(err, ErrEmptyStack) {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("removing last value, %v, using the pop method\n", v)
	}
	fmt.Println(s)

	s2 := Stack{}
	v, err = s2.Pop()
	if err != nil {
		// ignores if EmptyStackErr exists in err
		if !errors.Is(err, ErrEmptyStack) {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("removing last value, %v, using the pop method\n", v)
	}
	fmt.Println(s2)
}
