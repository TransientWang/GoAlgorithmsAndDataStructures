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

//53. 最大子序和
func maxSubArray(nums []int) int {
	var sum int
	res := nums[0]
	for _, val := range nums {
		if sum < 0 {
			sum = val
		} else {
			sum += val
		}
		if sum > res {
			res = sum
		}
	}
	return res
}
