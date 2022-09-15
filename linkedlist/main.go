package main

import "fmt"

// Node is a collection of two sub-elements or parts. A data part that stores the element and a next part that stores the link to the next node
type Node struct {
	data int
	next *Node
}

// LinkedList is formed when many such nodes are linked together to form a chain.
type LinkedList struct {
	head   *Node
	length int
}

// Prepend prepends a node to the LinkedList
func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// PrintListData prints all data values of the nodes of the LinkedList
func (l LinkedList) PrintListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

// DeleteWithValue deletes the node in the linkedList if the node with value of data is equal to the passed value to the method
func (l *LinkedList) DeleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--

}
func main() {
	myList := LinkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 18}
	node3 := &Node{data: 15}
	node4 := &Node{data: 5}
	node5 := &Node{data: 44}
	node6 := &Node{data: 77}
	node7 := &Node{data: 23}
	node8 := &Node{data: 7}
	myList.Prepend(node1)
	myList.Prepend(node2)
	myList.Prepend(node3)
	myList.Prepend(node4)
	myList.Prepend(node5)
	myList.Prepend(node6)
	myList.Prepend(node7)
	myList.Prepend(node8)
	myList.PrintListData()
	myList.DeleteWithValue(23)
	myList.PrintListData()
	myList.DeleteWithValue(235) // should be ignored
	myList.PrintListData()
}
