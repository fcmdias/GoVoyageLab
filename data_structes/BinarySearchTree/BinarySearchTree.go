package main

import "fmt"

// Node
type Node struct {
	Key int
	Left *Node
	Right *Node
}

// Insert
func (n *Node) Insert(v int) {
	if n.Key < v {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: v}
		} else {
			n.Right.Insert(v)
		}
	} else if n.Key > v {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: v}
		} else {
			n.Left.Insert(v)
		}
	}
}

// Search
func (n *Node) Search(v int) bool {
	if n == nil {
		return false
	}

	if n.Key < v {
		return n.Right.Search(v)
	} else if n.Key > v {
		return n.Left.Search(v)
	}
	// else v == n.Key
	return true
}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.Insert(4)
	tree.Insert(45)
	tree.Insert(64)
	tree.Insert(22)
	tree.Insert(220)
	tree.Insert(350)
	tree.Insert(270)
	tree.Insert(70)
	tree.Insert(11)

	fmt.Println(tree.Search(22))
	fmt.Println(tree.Search(23))

}