//go:build ignore

package main

import "fmt"

type Edge struct {
	to     int
	weight int
}

type WeightedGraph struct {
	vertices int
	adjList  [][]Edge
}

func NewWeightedGraph(vertices int) *WeightedGraph {
	return &WeightedGraph{
		vertices: vertices,
		adjList:  make([][]Edge, vertices),
	}
}

func (g *WeightedGraph) AddEdge(from, to, weight int) {
	g.adjList[from] = append(g.adjList[from], Edge{to: to, weight: weight})
}

func (g *WeightedGraph) AddUndirectedEdge(from, to, weight int) {
	g.AddEdge(from, to, weight)
	g.AddEdge(to, from, weight)
}

func (g *WeightedGraph) GetNeighbors(vertex int) []Edge {
	return g.adjList[vertex]
}

func (g *WeightedGraph) Print() {
	for i := 0; i < g.vertices; i++ {
		fmt.Printf("Vertex %d: ", i)
		for _, edge := range g.adjList[i] {
			fmt.Printf("(%d, %d) ", edge.to, edge.weight)
		}
		fmt.Println()
	}
}

func main() {
	graph := NewWeightedGraph(5)
	
	graph.AddUndirectedEdge(0, 1, 10)
	graph.AddUndirectedEdge(0, 4, 20)
	graph.AddUndirectedEdge(1, 2, 30)
	graph.AddUndirectedEdge(1, 3, 40)
	graph.AddUndirectedEdge(2, 3, 50)
	
	graph.Print()
}
