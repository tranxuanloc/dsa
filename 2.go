package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c := 0
	var prev *ListNode
	head := l2
	for l2 != nil {
		if l1 != nil {
			l2.Val += l1.Val
			l1 = l1.Next
		}
		l2.Val += c
		if l2.Val > 9 {
			l2.Val -= 10
			c = 1
		} else {
			c = 0
		}
		prev = l2
		l2 = l2.Next
	}
	prev.Next = l1
	for l1 != nil {
		l1.Val += c
		if l1.Val > 9 {
			l1.Val -= 10
			c = 1
		} else {
			c = 0
		}
		prev = l1
		l1 = l1.Next
	}
	if c > 0 {
		prev.Next = &ListNode{Val: c}
	}
	return head
}
