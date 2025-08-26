//go:build ignore

package main

import "fmt"

func printBoard(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
}

func isValid(board [][]int, row, col, num int) bool {
	// Check row
	for j := 0; j < 9; j++ {
		if board[row][j] == num {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// Check 3x3 box
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}

	return true
}

func findEmptyCell(board [][]int) []int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return []int{i, j}
			}
		}
	}
	return nil
}

func sudokuSolver(board [][]int) bool {
	emptyCell := findEmptyCell(board)
	if emptyCell == nil {
		return true
	}

	row, col := emptyCell[0], emptyCell[1]
	for num := 1; num <= 9; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if sudokuSolver(board) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}

func main() {
	board := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if sudokuSolver(board) {
		printBoard(board)
	} else {
		fmt.Println("No solution exists")
	}
}
