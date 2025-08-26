package main

import "fmt"

// Node represents a node in the binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert inserts a new value into the binary tree
func (n *Node) Insert(value int) {
	if n == nil {
		return
	}
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Search searches for a value in the binary tree
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}
	if value == n.Value {
		return true
	}
	if value < n.Value {
		return n.Left.Search(value)
	}
	return n.Right.Search(value)
}

// InOrderTraversal performs an in-order traversal of the binary tree
func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal()
	fmt.Print(n.Value, " ")
	n.Right.InOrderTraversal()
}

// main function to demonstrate the binary tree
func main() {
	root := &Node{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Print("In-order traversal: ")
	root.InOrderTraversal()
	fmt.Println()

	fmt.Println("Search 7:", root.Search(7))
	fmt.Println("Search 20:", root.Search(20))
}
