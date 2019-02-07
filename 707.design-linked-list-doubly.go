/*
 * @lc app=leetcode id=707 lang=golang
 *
 * [707] Design Linked List
 *
 * https://leetcode.com/problems/design-linked-list/description/
 *
 * algorithms
 * Easy (22.39%)
 * Total Accepted:    17.6K
 * Total Submissions: 78.3K
 * Testcase Example:  '["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]\n[[],[1],[3],[1,2],[1],[1],[1]]'
 *
 * Design your implementation of the linked list. You can choose to use the
 * singly linked list or the doubly linked list. A node in a singly linked list
 * should have two attributes: val and next. val is the value of the current
 * node, and next is a pointer/reference to the next node. If you want to use
 * the doubly linked list, you will need one more attribute prev to indicate
 * the previous node in the linked list. Assume all nodes in the linked list
 * are 0-indexed.
 *
 * Implement these functions in your linked list class:
 *
 *
 * get(index) : Get the value of the index-th node in the linked list. If the
 * index is invalid, return -1.
 * addAtHead(val) : Add a node of value val before the first element of the
 * linked list. After the insertion, the new node will be the first node of the
 * linked list.
 * addAtTail(val) : Append a node of value val to the last element of the
 * linked list.
 * addAtIndex(index, val) : Add a node of value val before the index-th node in
 * the linked list. If index equals to the length of linked list, the node will
 * be appended to the end of linked list. If index is greater than the length,
 * the node will not be inserted.
 * deleteAtIndex(index) : Delete the index-th node in the linked list, if the
 * index is valid.
 *
 *
 * Example:
 *
 *
 * MyLinkedList linkedList = new MyLinkedList();
 * linkedList.addAtHead(1);
 * linkedList.addAtTail(3);
 * linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
 * linkedList.get(1);            // returns 2
 * linkedList.deleteAtIndex(1);  // now the linked list is 1->3
 * linkedList.get(1);            // returns 3
 *
 *
 * Note:
 *
 *
 * All values will be in the range of [1, 1000].
 * The number of operations will be in the range of [1, 1000].
 * Please do not use the built-in LinkedList library.
 *
 *
 */
type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	prev *Node
	next *Node
	val  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	node := this.getNode(index)
	if node == nil {
		return -1
	}
	return node.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		val: val,
	}

	node := this.head
	if this.head == nil {
		this.head = newNode
		this.tail = newNode
	} else {
		newNode.next = this.head
		this.head = newNode
		node.prev = this.head
	}

	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{
		val: val,
	}

	if this.tail == nil {
		this.head = newNode
		this.tail = newNode
	} else {
		newNode.prev = this.tail
		this.tail.next = newNode
		this.tail = newNode
	}

	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	} else if index == this.size {
		this.AddAtTail(val)
		return
	}

	node := this.getNode(index)
	newNode := &Node{
		prev: node.prev,
		next: node,
		val:  val,
	}

	node.prev.next = newNode
	node.prev = newNode
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.size-1 {
		return
	}

	if index == 0 {
		this.head = this.head.next
		this.head.prev = nil
		this.size--
		return
	}

	if index == this.size-1 {
		this.tail = this.tail.prev
		this.tail.next = nil
		this.size--
		return
	}

	node := this.getNode(index)
	node.prev.next = node.next
	node.next.prev = node.prev
	this.size--
}

func (this *MyLinkedList) getNode(index int) *Node {
	if index > this.size-1 {
		return nil
	}

	var node *Node

	if index <= this.size/2 {
		// Iterate from head
		node = this.head
		for i := 0; i < index; i++ {
			node = node.next
		}
	} else {
		// Iterate from tail
		node = this.tail
		for i := this.size - 1; i > index; i-- {
			node = node.prev
		}
	}

	return node
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
