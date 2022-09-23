package main

import "fmt"

// Queue represents a queue that holds a slice
// like a queue of people, first in first out
type Queue struct {
	items []int
}

// Enqueue addas a value at the end
func (q *Queue) Enqueue(v int) {
	q.items = append(q.items, v)
}

// Dequeue removes the first value
func (q *Queue) Dequeue() int {
	var toBeRemoved int
	l := len(q.items)
	if l > 0 {
		toBeRemoved = q.items[0]
		q.items = q.items[1:]
	}
	return toBeRemoved
}

func main() {
	q := Queue{}
	q.Enqueue(3)
	q.Enqueue(7)
	q.Enqueue(2)
	q.Enqueue(9)
	fmt.Println(q)
	fmt.Printf("removing value, %v, from queue using method Dequeue\n", q.Dequeue())
	fmt.Println(q)
}
