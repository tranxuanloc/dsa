package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	out := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil {
		if cur.Left != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}
		out = append(out, cur.Val)
		cur = cur.Right
		if cur != nil {
			continue
		}
		for len(stack) > 0 && cur == nil {
			par := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			out = append(out, par.Val)
			cur = par.Right
		}
	}
	return out
}
