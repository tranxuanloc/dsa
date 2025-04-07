package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

func connect(root *Node116) *Node116 {
	queue := make([]*Node116, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		size := len(queue)
		for i := range size {
			node := queue[0]
			if i < size-1 {
				node.Next = queue[1]
			}
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
