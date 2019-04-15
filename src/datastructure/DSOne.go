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

//235. 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
