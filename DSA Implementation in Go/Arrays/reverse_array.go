//go:build ignore

package main

import "fmt"

func reverseArray(arr []int) []int {
	start := 0
	end := len(arr) - 1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	
	reversed := reverseArray(arr)
	fmt.Printf("Reversed array: %v\n", reversed)
}
