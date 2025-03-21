package main

type MyLinkedList struct {
	size int
	head *Node
	tail *Node
}

type Node struct {
	val  int
	prev *Node
	next *Node
}

func Constructor() MyLinkedList {
	head := &Node{}
	tail := &Node{prev: head}
	head.next = tail
	return MyLinkedList{head: head, tail: tail}
}

func (this *MyLinkedList) Get(index int) int {
	node := this.getNode(index)
	if node == nil {
		return -1
	}
	return node.val
}

func (this *MyLinkedList) getNode(index int) *Node {
	if index < 0 || index > this.size-1 {
		return nil
	}
	if step := this.size - index - 1; step < index {
		cur := this.tail.prev
		for ; step > 0; step-- {
			cur = cur.prev
		}
		return cur
	} else {
		cur := this.head.next
		for ; index > 0; index-- {
			cur = cur.next
		}
		return cur
	}
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{val: val, next: this.head.next, prev: this.head}
	this.head.next.prev = newNode
	this.head.next = newNode
	this.size = this.size + 1
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{val: val, next: this.tail, prev: this.tail.prev}
	this.tail.prev.next = newNode
	this.tail.prev = newNode
	this.size = this.size + 1
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.size {
		this.AddAtTail(val)
	} else if index == 0 {
		this.AddAtHead(val)
	} else if node := this.getNode(index); node != nil {
		newNode := &Node{val: val, next: node, prev: node.prev}
		node.prev.next = newNode
		node.prev = newNode
		this.size = this.size + 1
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	node := this.getNode(index)
	if node != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
		this.size = this.size - 1
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
