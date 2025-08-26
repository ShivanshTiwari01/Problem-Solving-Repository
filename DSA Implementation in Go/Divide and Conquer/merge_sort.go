package main

import "fmt"

// MergeSort sorts an array using the merge sort algorithm
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// merge merges two sorted arrays into one sorted array
func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	// Append any remaining elements from either array
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}

// main function to demonstrate merge sort
func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Print("Unsorted array: ")
	fmt.Println(arr)

	sorted := MergeSort(arr)
	fmt.Print("Sorted array: ")
	fmt.Println(sorted)
}
