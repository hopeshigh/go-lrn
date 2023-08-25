package ggggraphfundamentals

import (
	"reflect"
	"testing"
)

func TestAdjacencyList(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)

	tests := []struct {
		describe string
		input    int
		expected []int
	}{
		{"first graph index", 1, []int{2, 3}},
		{"second graph index", 2, []int{1, 4}},
		{"invalid graph index", 5, nil},
	}

	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			got := g.AdjList[tt.input]
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("input: %v, got: %v, expected: %v", tt.input, got, tt.expected)
			}
		})
	}
}
