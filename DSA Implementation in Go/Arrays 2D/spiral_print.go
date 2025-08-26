//go:build ignore

package main

import "fmt"

func spiralPrint(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// Print top row
		for i := left; i <= right; i++ {
			fmt.Print(matrix[top][i], " ")
		}
		top++

		// Print right column
		for i := top; i <= bottom; i++ {
			fmt.Print(matrix[i][right], " ")
		}
		right--

		// Print bottom row
		if top <= bottom {
			for i := right; i >= left; i-- {
				fmt.Print(matrix[bottom][i], " ")
			}
			bottom--
		}

		// Print left column
		if left <= right {
			for i := bottom; i >= top; i-- {
				fmt.Print(matrix[i][left], " ")
			}
			left++
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	spiralPrint(matrix)
}
