//go:build ignore

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

type Queue struct {
	stack1 *Stack
	stack2 *Stack
}

func (q *Queue) Enqueue(data int) {
	q.stack1.push(data)
}

func (q *Queue) Dequeue() int {
	if q.stack2.isEmpty() {
		for !q.stack1.isEmpty() {
			q.stack2.push(q.stack1.pop())
		}
	}
	return q.stack2.pop()
}

func main() {
	queue := &Queue{
		stack1: &Stack{},
		stack2: &Stack{},
	}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Dequeued element:", queue.Dequeue())
	fmt.Println("Dequeued element:", queue.Dequeue())
}
