package main

import "fmt"

type MinHeap struct {
	elements []int
}

func (h *MinHeap) Insert(value int) {
	h.elements = append(h.elements, value)
	h.heapifyUp(len(h.elements) - 1)
}

func (h *MinHeap) ExtractMin() int {
	if len(h.elements) == 0 {
		return -1 // or panic("heap is empty")
	}
	
	min := h.elements[0]
	lastIdx := len(h.elements) - 1
	h.elements[0] = h.elements[lastIdx]
	h.elements = h.elements[:lastIdx]
	
	if len(h.elements) > 0 {
		h.heapifyDown(0)
	}
	
	return min
}

func (h *MinHeap) heapifyUp(index int) {
	parent := (index - 1) / 2
	if index > 0 && h.elements[index] < h.elements[parent] {
		h.elements[index], h.elements[parent] = h.elements[parent], h.elements[index]
		h.heapifyUp(parent)
	}
}

func (h *MinHeap) heapifyDown(index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	smallest := index
	
	if leftChild < len(h.elements) && h.elements[leftChild] < h.elements[smallest] {
		smallest = leftChild
	}
	
	if rightChild < len(h.elements) && h.elements[rightChild] < h.elements[smallest] {
		smallest = rightChild
	}
	
	if smallest != index {
		h.elements[index], h.elements[smallest] = h.elements[smallest], h.elements[index]
		h.heapifyDown(smallest)
	}
}

func (h *MinHeap) Peek() int {
	if len(h.elements) == 0 {
		return -1
	}
	return h.elements[0]
}

func (h *MinHeap) Size() int {
	return len(h.elements)
}

func main() {
	heap := &MinHeap{}
	
	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(15)
	heap.Insert(3)
	heap.Insert(8)
	
	fmt.Println("Heap size:", heap.Size())
	fmt.Println("Min element:", heap.Peek())
	
	for heap.Size() > 0 {
		fmt.Println("Extracted:", heap.ExtractMin())
	}
}
