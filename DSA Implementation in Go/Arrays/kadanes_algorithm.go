//go:build ignore

package main

import "fmt"

func kadanesAlgorithm(arr []int) int {
	maxSoFar := arr[0]
	currentMax := arr[0]

	for i := 1; i < len(arr); i++ {
		currentMax = max(arr[i], currentMax+arr[i])
		maxSoFar = max(maxSoFar, currentMax)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := kadanesAlgorithm(arr)
	
	fmt.Printf("Maximum subarray sum is: %d\n", result)
}
