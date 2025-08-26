package main

import (
	"fmt"
)

type HashTable struct {
	buckets [][]KeyValuePair
	size    int
}

type KeyValuePair struct {
	key   string
	value string
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([][]KeyValuePair, size),
		size:    size,
	}
}

func (h *HashTable) Put(key, value string) {
	index := hash(key) % h.size
	h.buckets[index] = append(h.buckets[index], KeyValuePair{key: key, value: value})
}

func (h *HashTable) Get(key string) (string, bool) {
	index := hash(key) % h.size
	for _, pair := range h.buckets[index] {
		if pair.key == key {
			return pair.value, true
		}
	}
	return "", false
}

func (h *HashTable) Remove(key string) {
	index := hash(key) % h.size
	for i, pair := range h.buckets[index] {
		if pair.key == key {
			h.buckets[index] = append(h.buckets[index][:i], h.buckets[index][i+1:]...)
			return
		}
	}
}

func hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash
}

func main() {
	hashTable := NewHashTable(10)
	hashTable.Put("name", "Alice")
	hashTable.Put("age", "25")

	if value, found := hashTable.Get("name"); found {
		fmt.Println("Name:", value)
	}

	hashTable.Remove("age")
	if _, found := hashTable.Get("age"); !found {
		fmt.Println("Age not found")
	}
}
