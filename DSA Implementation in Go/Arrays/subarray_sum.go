//go:build ignore

package main

import "fmt"

func subarraySum(arr []int, target int) (int, int) {
	currentSum := 0
	start := 0

	for end := 0; end < len(arr); end++ {
		currentSum += arr[end]

		for currentSum > target && start <= end {
			currentSum -= arr[start]
			start++
		}

		if currentSum == target {
			return start, end
		}
	}

	return -1, -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 9
	start, end := subarraySum(arr, target)

	if start != -1 {
		fmt.Printf("Subarray found between indices %d and %d\n", start, end)
	} else {
		fmt.Println("No subarray found")
	}
}
