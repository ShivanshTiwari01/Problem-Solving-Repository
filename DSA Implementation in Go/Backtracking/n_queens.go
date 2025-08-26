//go:build ignore

package main

import "fmt"

func printBoard(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 1 {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func isSafe(board [][]int, row, col int) bool {
	// Check this column on upper side
	for i := 0; i < row; i++ {
		if board[i][col] == 1 {
			return false
		}
	}

	// Check upper diagonal on left side
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	// Check upper diagonal on right side
	for i, j := row, col; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}

func solveNQueensUtil(board [][]int, row int) bool {
	if row == len(board) {
		printBoard(board)
		return true
	}

	for col := 0; col < len(board); col++ {
		if isSafe(board, row, col) {
			board[row][col] = 1
			if solveNQueensUtil(board, row+1) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}

func solveNQueens(n int) {
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}
	if !solveNQueensUtil(board, 0) {
		fmt.Println("No solution exists")
	}
}

func main() {
	n := 4
	solveNQueens(n)
}
