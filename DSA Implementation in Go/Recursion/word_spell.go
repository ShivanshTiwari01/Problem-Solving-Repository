//go:build ignore

package main

import (
	"fmt"
)

func spellWord(word string) {
	if len(word) == 0 {
		return
	}
	fmt.Println(word)
	spellWord(word[1:])
}

func main() {
	word := "hello"
	spellWord(word)
}
