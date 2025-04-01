package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	order := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			order = append(order, cur.Val)
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			cur = cur.Left
			continue
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return order
}
