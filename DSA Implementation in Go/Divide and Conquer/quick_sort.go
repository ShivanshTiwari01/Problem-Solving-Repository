//go:build ignore

package main

import "fmt"

// QuickSort sorts an array using the quick sort algorithm
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}

	for i := 0; i < len(arr); i++ {
		if i == len(arr)/2 {
			continue
		}
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)
}

// main function to demonstrate quick sort
func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Print("Unsorted array: ")
	fmt.Println(arr)

	sorted := QuickSort(arr)
	fmt.Print("Sorted array: ")
	fmt.Println(sorted)
}
