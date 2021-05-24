package datastructure

import "math"

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

func avlInsert(root *TreeNode, val int) {
	x := root
	var (
		y       *TreeNode
		newNode = &TreeNode{
			Val: val,
		}
	)
	for x != nil {
		y = x
		if val < x.Val {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	if y == nil {
		root = newNode
	} else if val < y.Val {
		y.Left = newNode
	} else {
		y.Right = newNode
	}
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(float64(maxDepth(root.Left)-maxDepth(root.Right))) <= 1 && isBalanced(root.Right) && isBalanced(root.Left)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func isbl(a, b int) bool {
	r := a - b
	return -1 <= r && r <= 1
}
func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(maxDepth(node.Right)+1, maxDepth(node.Left)+1)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return left == nil || right == nil
	}
	if left.Val != right.Val{
		return false
	}
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)

}
