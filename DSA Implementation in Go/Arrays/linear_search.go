//go:build ignore

package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{5, 3, 4, 1, 2}
	target := 4

	result := linearSearch(arr, target)

	if result != -1 {
		fmt.Printf("Element found at index: %d\n", result)
	} else {
		fmt.Println("Element not found")
	}

}
