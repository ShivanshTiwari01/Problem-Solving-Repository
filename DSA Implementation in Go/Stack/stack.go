package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) push(data int) {
	newNode := &Node{data: data}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) pop() int {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return -1
	}
	popped := s.top.data
	s.top = s.top.next
	return popped
}

func (s *Stack) peek() int {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.top.data
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}

func main() {
	stack := &Stack{}
	stack.push(10)
	stack.push(20)
	stack.push(30)

	fmt.Println("Top element:", stack.peek())
	fmt.Println("Popped element:", stack.pop())
	fmt.Println("Is stack empty?", stack.isEmpty())
}
