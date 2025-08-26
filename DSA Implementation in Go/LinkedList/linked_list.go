package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the linked list structure
type LinkedList struct {
	head *Node
}

// Insert adds a new node at the beginning of the list
func (ll *LinkedList) Insert(data int) {
	newNode := &Node{data: data, next: ll.head}
	ll.head = newNode
}

// Delete removes the first occurrence of the specified data
func (ll *LinkedList) Delete(data int) {
	if ll.head == nil {
		return
	}
	
	if ll.head.data == data {
		ll.head = ll.head.next
		return
	}
	
	current := ll.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}
	
	if current.next != nil {
		current.next = current.next.next
	}
}

// Display prints all elements in the list
func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// Search finds if a value exists in the list
func (ll *LinkedList) Search(data int) bool {
	current := ll.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

// Size returns the number of nodes in the list
func (ll *LinkedList) Size() int {
	count := 0
	current := ll.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func main() {
	ll := &LinkedList{}
	
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	
	fmt.Print("List: ")
	ll.Display()
	
	fmt.Printf("Size: %d\n", ll.Size())
	fmt.Printf("Search 20: %t\n", ll.Search(20))
	
	ll.Delete(20)
	fmt.Print("After deleting 20: ")
	ll.Display()
}
