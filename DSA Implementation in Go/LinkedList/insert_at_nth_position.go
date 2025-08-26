//go:build ignore

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) insertAtNth(data, position int) {
	newNode := &Node{data: data}
	
	if position == 0 {
		newNode.next = ll.head
		ll.head = newNode
		return
	}
	
	current := ll.head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.next
	}
	
	if current == nil {
		fmt.Println("Position out of bounds")
		return
	}
	
	newNode.next = current.next
	current.next = newNode
}

func (ll *LinkedList) display() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := &LinkedList{}
	
	ll.insertAtNth(1, 0)
	ll.insertAtNth(3, 1)
	ll.insertAtNth(5, 2)
	ll.insertAtNth(2, 1)
	
	fmt.Print("Linked List: ")
	ll.display()
}
