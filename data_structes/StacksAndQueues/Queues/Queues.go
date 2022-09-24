package main

import (
	"errors"
	"fmt"
)

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
func (q *Queue) Dequeue() (int, error) {
	var toBeRemoved int
	l := len(q.items)
	if l > 0 {
		toBeRemoved = q.items[0]
		q.items = q.items[1:]
		return toBeRemoved, nil
	}
	return 0, errors.New("Queue is empty therefore cannot remove value using the Dequeue method")

}

func main() {
	q := Queue{}
	q.Enqueue(3)
	q.Enqueue(7)
	q.Enqueue(2)
	q.Enqueue(9)
	fmt.Println(q)
	i, err := q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("removing value, %v, from queue using method Dequeue\n", i)
	}
	fmt.Println(q)

	q2 := Queue{}
	i, err = q2.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("removing value, %v, from queue using method Dequeue\n", i)
	}
	fmt.Println(q2)
}
