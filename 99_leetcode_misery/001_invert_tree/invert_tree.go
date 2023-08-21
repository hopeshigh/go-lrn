package inverttree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	// Base case
	if root == nil {
		return nil
	}

	// Invert left and right subtrees
	leftInverted := InvertTree(root.Left)
	rightInverted := InvertTree(root.Right)

	// Swap left and right subtrees
	root.Left = rightInverted
	root.Right = leftInverted

	return root
}
