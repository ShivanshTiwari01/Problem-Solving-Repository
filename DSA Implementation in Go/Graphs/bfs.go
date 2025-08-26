package main

import "fmt"

type Graph struct {
	vertices int
	adjList  [][]int
}

func NewGraph(v int) *Graph {
	return &Graph{
		vertices: v,
		adjList:  make([][]int, v),
	}
}

func (g *Graph) AddEdge(src, dest int) {
	g.adjList[src] = append(g.adjList[src], dest)
	g.adjList[dest] = append(g.adjList[dest], src) // for undirected graph
}

func (g *Graph) BFS(start int) {
	visited := make([]bool, g.vertices)
	queue := []int{start}
	visited[start] = true

	fmt.Printf("BFS traversal starting from vertex %d: ", start)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", current)

		for _, neighbor := range g.adjList[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	fmt.Println()
}

func main() {
	g := NewGraph(6)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)
	g.AddEdge(4, 5)

	g.BFS(0)
}
