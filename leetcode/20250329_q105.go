package main

// 先序遍历和中序遍历可以确定一个二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var res = &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return res
	}
	// 获得根节点
	var root, index = preorder[0], 0
	// 区分开左分支和右分支
	// 找到根节点在中序的位置并进行分开左分支和右分支
	for i := 0; i < len(inorder); i++ {
		if root == inorder[i] {
			index = i
			break
		}
	}
	inLeft := inorder[0:index]
	inRight := inorder[index+1:]
	preleft := preorder[1 : len(inLeft)+1]
	preRight := preorder[len(inLeft)+1:]
	res.Left = buildTree(preleft, inLeft)
	res.Right = buildTree(preRight, inRight)
	return res
}
