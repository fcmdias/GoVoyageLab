package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an elements to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() {
	if len(h.array) == 0 {
		return
	}
	l := len(h.array) - 1
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyUp(0)

}

// maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify from bottom top
func (h *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {

		if l == lastIndex { // when index is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
	}

	// compare array value of current index to larger child and swap if smaller
	if h.array[index] < h.array[childToCompare] {
		h.swap(index, childToCompare)
		index = childToCompare
		l, r = left(index), right(index)
	}
}

// parent returns the parent's index
func parent(i int) int {
	return (i - 1) / 2
}

// left returns the left child index
func left(i int) int {
	return 2*1 + 1
}

// right returns the left child index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 54, 2, 6, 33, 88, 100, 23}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 3; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
