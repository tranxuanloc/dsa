package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var prev *TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}
		cur = stack[len(stack)-1]
		if cur.Right == nil || cur.Right == prev {
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			prev = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
	return res
}

func traverse(root *TreeNode, res *[]int) {
	if root != nil {
		traverse(root.Left, res)
		traverse(root.Right, res)
		*res = append(*res, root.Val)
	}
}
