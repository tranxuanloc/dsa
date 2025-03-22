package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	odd := head
	if head == nil {
		return nil
	}
	even := head.Next
	for even != nil {
		temp := even.Next
		if temp != nil {
			even.Next = temp.Next
			temp.Next = odd.Next
			odd.Next = temp
		}
		odd = temp
		even = even.Next
	}
	return head
}
