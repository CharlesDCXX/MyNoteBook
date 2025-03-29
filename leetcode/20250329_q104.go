package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return deepth(root, 1)
}
func deepth(root *TreeNode, high int) int {
	var leftDeep, rightDeep = high, high
	if root.Left != nil {
		leftDeep = deepth(root.Left, leftDeep+1)
	}
	if root.Right != nil {
		rightDeep = deepth(root.Right, rightDeep+1)
	}
	return max(leftDeep, rightDeep)
}

// 广度
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return deep([]*TreeNode{root}, 1)

}

func deep(root []*TreeNode, high int) int {
	var trees []*TreeNode
	for _, node := range root {
		if node.Left != nil {
			trees = append(trees, node.Left)
		}
		if node.Right != nil {
			trees = append(trees, node.Right)
		}
	}
	if len(trees) == 0 {
		return high
	}
	return deep(trees, high+1)
}
