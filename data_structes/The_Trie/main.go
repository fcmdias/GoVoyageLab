package main

import "fmt"

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents  each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the roor node
type Trie struct {
	root *Node
}

// InitTrie iniciates a trie
func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {

	currentNode := t.root
	for i := 0; i < len(w); i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

// Search will take in a word and return true if that word is included in the trie
func (t *Trie) Search(w string) bool {
	currentNode := t.root
	for i := 0; i < len(w); i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if !currentNode.isEnd {
		return false
	}
	return true

}

func main() {
	t := InitTrie()
	t.Insert("car")
	t.Insert("cart")
	t.Insert("casa")
	t.Insert("water")
	t.Insert("walter")
	t.Insert("wasp")
	t.Insert("space")
	t.Insert("spar")
	t.Insert("spoon")

	fmt.Println(t.Search("spin"))
	fmt.Println(t.Search("space"))

}
