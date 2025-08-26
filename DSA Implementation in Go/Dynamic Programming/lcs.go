package main

import "fmt"

// LCS computes the length of the longest common subsequence of two strings
func LCS(X, Y string) int {
	m := len(X)
	n := len(Y)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// main function to demonstrate LCS
func main() {
	X := "AGGTAB"
	Y := "GXTXAYB"
	fmt.Print("Length of LCS: ")
	fmt.Println(LCS(X, Y))
}
