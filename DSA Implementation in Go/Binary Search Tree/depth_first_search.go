//go:build ignore

package main

import "fmt"

// Node represents a node in the binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// DFS performs a depth-first search on the binary tree
func (n *Node) DFS() {
	if n == nil {
		return
	}
	fmt.Print(n.Value, " ")
	n.Left.DFS()
	n.Right.DFS()
}

// main function to demonstrate DFS
func main() {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	fmt.Print("Depth First Search: ")
	root.DFS()
	fmt.Println()
}
