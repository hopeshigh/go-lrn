package ggggraphfundamentals

import "fmt"

// Prefer Adjacency Lists if the graph is "sparse". "Sparse" can be defined by
// having a lot of blank or empty connections

// todo: learn how to make this use "generics"
type ALGraph struct {
	AdjList map[int][]int
}

func NewGraph() *ALGraph {
	return &ALGraph{AdjList: make(map[int][]int)}
}

func (g *ALGraph) AddEdge(src int, dest int) {
	g.AdjList[src] = append(g.AdjList[src], dest)

	// For an undirected graph, also add an edge from dest to src
	// Small note: The AdjList key changes here
	g.AdjList[dest] = append(g.AdjList[dest], src)
}

func (g *ALGraph) Print() {
	for vertex, edges := range g.AdjList {
		fmt.Printf("Vertex %v: ", vertex)
		for _, edge := range edges {
			fmt.Printf("%v ", edge)
		}
		fmt.Println()
	}
}
