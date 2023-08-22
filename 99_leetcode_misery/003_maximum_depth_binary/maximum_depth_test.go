package maximumdepth

import (
	"testing"
)

func TestRecurMaximumDepth(t *testing.T) {
	tests := []struct {
		desc     string
		input    *TreeNode
		expected int
	}{
		{"sample tree",
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			}, 3},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got := RecurMaximumDepth(test.input)
			if got != test.expected {
				t.Errorf("want: %v, got: %v", test.expected, got)
			}
		})
	}
}

func TestBFSMaximumDepth(t *testing.T) {
	tests := []struct {
		desc     string
		input    *TreeNode
		expected int
	}{
		{"sample tree",
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			}, 3},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got := BFSMaximumDepth(test.input)
			if got != test.expected {
				t.Errorf("want: %v, got: %v", test.expected, got)
			}
		})
	}
}
