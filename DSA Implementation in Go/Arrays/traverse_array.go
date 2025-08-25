package main

import "fmt"

func main() {

	var n int

	fmt.Print("Enter the size of the array:")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Print("Enter element ", i+1, ": ")
		fmt.Scan(&arr[i])
	}

	for i := range arr {
		fmt.Println(arr[i])
	}

}
