//go:build ignore

package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	
	return len(stack) == 0
}

func main() {
	testCases := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	
	for _, test := range testCases {
		result := isValid(test)
		fmt.Printf("Input: %s, Valid: %t\n", test, result)
	}
}
