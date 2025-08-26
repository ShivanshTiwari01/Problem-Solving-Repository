//go:build ignore

package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 4

	result := binarySearch(arr, target)

	if result != -1 {
		fmt.Printf("Element found at index: %d\n", result)
	} else {
		fmt.Println("Element not found")
	}
	
}
