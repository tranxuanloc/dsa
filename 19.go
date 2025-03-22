package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sz := 0
	c := head
	for c != nil {
		sz++
		c = c.Next
	}
	if n == sz {
		return head.Next
	}
	c = head
	for i := sz - n; i > 1; i-- {
		c = c.Next
	}
	c.Next = c.Next.Next
	return head
}
