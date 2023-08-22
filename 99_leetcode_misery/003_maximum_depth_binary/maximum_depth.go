package maximumdepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RecurMaximumDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := recurTree(root.Left, 1)
	right := recurTree(root.Right, 1)

	return max(left, right)
}

func recurTree(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	left := recurTree(node.Left, depth+1)
	right := recurTree(node.Right, depth+1)

	return max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func BFSMaximumDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}

			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
		depth++
	}
	return depth
}
