//go:build ignore

package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

type Stack struct {
	q1 Queue
	q2 Queue
}

func (s *Stack) Push(item int) {
	s.q2.Enqueue(item)
	for !s.q1.IsEmpty() {
		s.q2.Enqueue(s.q1.Dequeue())
	}
	s.q1, s.q2 = s.q2, s.q1
}

func (s *Stack) Pop() int {
	if s.q1.IsEmpty() {
		return -1
	}
	return s.q1.Dequeue()
}

func (s *Stack) Top() int {
	if s.q1.IsEmpty() {
		return -1
	}
	return s.q1.items[0]
}

func (s *Stack) IsEmpty() bool {
	return s.q1.IsEmpty()
}

func main() {
	stack := Stack{}
	
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	fmt.Println("Top:", stack.Top())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Top:", stack.Top())
	fmt.Println("IsEmpty:", stack.IsEmpty())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("IsEmpty:", stack.IsEmpty())
}
