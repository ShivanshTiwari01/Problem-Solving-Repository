//go:build ignore

package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents a weighted edge in the graph
type Edge struct {
	to     int
	weight int
}

// Graph represents an adjacency list graph
type Graph struct {
	vertices int
	adjList  [][]Edge
}

// NewGraph creates a new graph with given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make([][]Edge, vertices),
	}
}

// AddEdge adds a weighted edge to the graph
func (g *Graph) AddEdge(from, to, weight int) {
	g.adjList[from] = append(g.adjList[from], Edge{to: to, weight: weight})
}

// Item represents an item in the priority queue
type Item struct {
	vertex   int
	distance int
	index    int
}

// PriorityQueue implements heap.Interface for Dijkstra's algorithm
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Dijkstra finds shortest paths from source vertex to all other vertices
func (g *Graph) Dijkstra(source int) []int {
	dist := make([]int, g.vertices)
	visited := make([]bool, g.vertices)
	
	// Initialize distances to infinity
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[source] = 0
	
	// Create priority queue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{vertex: source, distance: 0})
	
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)
		u := current.vertex
		
		if visited[u] {
			continue
		}
		visited[u] = true
		
		// Check all adjacent vertices
		for _, edge := range g.adjList[u] {
			v := edge.to
			weight := edge.weight
			
			if !visited[v] && dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(&pq, &Item{vertex: v, distance: dist[v]})
			}
		}
	}
	
	return dist
}

func main() {
	// Example usage
	g := NewGraph(6)
	
	// Add edges (from, to, weight)
	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 2, 2)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 8)
	g.AddEdge(2, 4, 10)
	g.AddEdge(3, 4, 2)
	g.AddEdge(3, 5, 6)
	g.AddEdge(4, 5, 3)
	
	source := 0
	distances := g.Dijkstra(source)
	
	fmt.Printf("Shortest distances from vertex %d:\n", source)
	for i, dist := range distances {
		if dist == math.MaxInt32 {
			fmt.Printf("Vertex %d: unreachable\n", i)
		} else {
			fmt.Printf("Vertex %d: %d\n", i, dist)
		}
	}
}
