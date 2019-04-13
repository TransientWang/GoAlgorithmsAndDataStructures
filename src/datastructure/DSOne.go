package datastructure

//Definition for a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

//226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftNode, rightNode := (*root).Left, root.Right
	invertTree(leftNode)
	invertTree(rightNode)
	root.Right, root.Left = leftNode, rightNode
	return root
}
