package main

import "fmt"

// IsPowerOfTwo checks if a number is a power of 2
func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// main function to demonstrate the power of 2 check
func main() {
	num := 16
	if IsPowerOfTwo(num) {
		fmt.Println(num, "is a power of 2.")
	} else {
		fmt.Println(num, "is not a power of 2.")
	}
}
