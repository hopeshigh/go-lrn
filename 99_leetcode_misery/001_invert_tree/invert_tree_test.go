package inverttree

import (
	"testing"
)

// Helper function to check if two trees are equal
func areEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && areEqual(t1.Left, t2.Left) && areEqual(t1.Right, t2.Right)
}

func TestInvertTree(t *testing.T) {
	// Original tree:
	//     1
	//    / \
	//   2   3
	//  / \
	// 4   5
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	// Expected inverted tree:
	//     1
	//    / \
	//   3   2
	//        / \
	//       5   4
	expected := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 4},
		},
	}

	invertedRoot := InvertTree(root)
	if !areEqual(invertedRoot, expected) {
		t.Errorf("Tree is not inverted. got: %v want: %v", invertedRoot, expected)
	}
}
