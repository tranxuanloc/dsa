package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	curCount := 0
	next := 0
	if root != nil {
		curCount = 1
		queue = append(queue, root)
	}
	var cur *TreeNode
	level := make([]int, 0)
	for len(queue) > 0 {
		cur = queue[0]
		level = append(level, cur.Val)
		curCount--
		queue = queue[1:]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			next++
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			next++
		}
		if curCount == 0 {
			res = append(res, level)
			level = make([]int, 0)
			curCount = next
			next = 0
		}
	}
	return res
}
