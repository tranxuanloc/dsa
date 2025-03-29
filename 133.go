package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node133 struct {
	Val       int
	Neighbors []*Node133
}

func cloneGraph(node *Node133) *Node133 {
	if node == nil {
		return nil
	}
	cache := make(map[int]*Node133)
	return clone(node, cache)
}

func clone(node *Node133, cache map[int]*Node133) *Node133 {
	if node == nil {
		return nil
	}
	copied := cache[node.Val]
	if copied != nil {
		return copied
	}
	copiedNeis := make([]*Node133, len(node.Neighbors))
	copied = &Node133{Val: node.Val, Neighbors: copiedNeis}
	cache[node.Val] = copied
	for i, nei := range node.Neighbors {
		copiedNeis[i] = clone(nei, cache)
	}
	return copied
}
