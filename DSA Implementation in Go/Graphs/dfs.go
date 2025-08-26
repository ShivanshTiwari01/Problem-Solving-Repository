//go:build ignore

package main

import "fmt"

type Graph struct {
	vertices int
	adjList  [][]int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make([][]int, vertices),
	}
}

func (g *Graph) AddEdge(src, dest int) {
	g.adjList[src] = append(g.adjList[src], dest)
}

func (g *Graph) DFS(startVertex int) {
	visited := make([]bool, g.vertices)
	g.dfsUtil(startVertex, visited)
}

func (g *Graph) dfsUtil(vertex int, visited []bool) {
	visited[vertex] = true
	fmt.Printf("%d ", vertex)
	
	for _, neighbor := range g.adjList[vertex] {
		if !visited[neighbor] {
			g.dfsUtil(neighbor, visited)
		}
	}
}

func main() {
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	
	fmt.Print("DFS traversal starting from vertex 0: ")
	g.DFS(0)
	fmt.Println()
}
