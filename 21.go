package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	node := &ListNode{}
	prev := node
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			prev.Next = list2
			list2 = list2.Next
		} else {
			prev.Next = list1
			list1 = list1.Next
		}
		prev = prev.Next
	}
	if list1 != nil {
		prev.Next = list1
	} else {
		prev.Next = list2
	}
	return node.Next
}
