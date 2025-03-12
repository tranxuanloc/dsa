package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	count := 1
	queue := make([]*TreeNode, 0)
	for root != nil {
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right == nil {
			break
		}
		level = level + 1
		count += int(math.Pow(2, float64(level)))
		root = root.Right
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
	}
	return count + len(queue)
}
