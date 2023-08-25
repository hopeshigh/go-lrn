package ggggraphfundamentals

import "fmt"

// Takes up more space (O(V^2)) but can be faster for certain operations
type AMGraph struct {
	AdjMatrix [][]int
}

func NewAMGraph(vertices int) *AMGraph {
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices)
	}

	return &AMGraph{AdjMatrix: matrix}
}

func (g *AMGraph) AddEdge(src int, dest int) {
	g.AdjMatrix[src][dest] = 1

	// For undirected graph, add an edge from dest to src
	g.AdjMatrix[dest][src] = 1
}

func (g *AMGraph) Print() {
	for i := 0; i < len(g.AdjMatrix); i++ {
		fmt.Println(g.AdjMatrix[i])
	}
}
