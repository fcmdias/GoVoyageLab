package main

import "fmt"

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
func (s *Stack) Pop() int {
	l := len(s.items)
	toRemove := s.items[l-1]
	s.items = s.items[:l-1]
	return toRemove
}

func main() {
	s := Stack{}
	s.Push(7)
	s.Push(9)
	s.Push(4)
	s.Push(8)
	fmt.Println(s)
	fmt.Printf("removing last value, %v, using the pop method\n", s.Pop())
	fmt.Println(s)

}
