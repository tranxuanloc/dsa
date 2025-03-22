package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := 0
	nodeA := headA
	for nodeA != nil {
		m++
		if nodeA.Next == nil {
			break
		}
		nodeA = nodeA.Next
	}
	n := 0
	nodeB := headB
	for nodeB != nil {
		n++
		if nodeB.Next == nil {
			break
		}
		nodeB = nodeB.Next
	}
	if nodeA != nodeB {
		return nil
	}
	k := n - m
	nodeA = headA
	nodeB = headB
	if m > n {
		nodeA, nodeB = nodeB, nodeA
		k = m - n
	}
	// fmt.Println(k, m, n)
	for i := k; i > 0; i-- {
		nodeB = nodeB.Next
	}
	for i := 0; i < m; i++ {
		if nodeA == nodeB {
			return nodeA
		}
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}
	return nil
}
