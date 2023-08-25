package ggggraphfundamentals

import (
	"reflect"
	"testing"
)

func TestAdjacencyMatrix(t *testing.T) {
	g := NewAMGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)

	tests := []struct {
		describe string
		input    []int
		expected int
	}{
		{"first graph index", []int{1, 2}, 1},
		{"second graph index", []int{1, 4}, 0},
		{"third graph index", []int{2, 4}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			got := g.AdjMatrix[tt.input[0]][tt.input[1]]
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("input: %v, got: %v, expected: %v", tt.input, got, tt.expected)
			}
		})
	}
}
