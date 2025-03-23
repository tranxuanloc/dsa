package main

import "runtime"

/**
 * Definition for a Node.
 */
type Node2 struct {
	Val   int
	Prev  *Node2
	Next  *Node2
	Child *Node2
}

func flatten(root *Node2) *Node2 {
	defer runtime.GC()
	stack := make([]*Node2, 0)
	p := -1
	head := root
	cur := head
	for cur != nil {
		if cur.Child != nil {
			stack = append(stack, cur)
			p++
			cur = cur.Child
		} else {
			if cur.Next == nil && p > -1 {
				par := stack[p]
				cur.Next = par.Next
				if par.Next != nil {
					par.Next.Prev = cur
				}
				par.Next = par.Child
				par.Child.Prev = par
				par.Child = nil
				p--
			}
			if cur.Next != nil || p < 0 {
				cur = cur.Next
			}
		}
	}
	return head
}
