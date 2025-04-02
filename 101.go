package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return helper(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
	if left == right {
		return true
	}
	if left == nil && right != nil || left != nil && right == nil || left.Val != right.Val {
		return false
	}
	isLeft := helper(left.Left, right.Right)
	isRight := helper(left.Right, right.Left)
	return isLeft && isRight
}
