package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (t *Trie) Insert(word string) {
	current := t.root
	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			current.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		current = current.children[char]
	}
	current.isEnd = true
}

func (t *Trie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}
	return current.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t.root
	for _, char := range prefix {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}
	return true
}

func main() {
	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("app")
	
	fmt.Println(trie.Search("app"))     // true
	fmt.Println(trie.Search("apple"))   // true
	fmt.Println(trie.Search("appl"))    // false
	fmt.Println(trie.StartsWith("app")) // true
}
