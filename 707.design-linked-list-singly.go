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
}

type Node struct {
	next *Node
	val  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	node := this.head

	for i := 0; i < index; i++ {
		if node == nil {
			return -1
		}
		node = node.next
	}

	if node == nil {
		return -1
	}
	return node.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &Node{
		next: this.head,
		val:  val,
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.head = &Node{
			next: nil,
			val:  val,
		}
		return
	}

	node := this.head

	for {
		if node.next == nil {
			node.next = &Node{
				next: nil,
				val:  val,
			}
			return
		}
		node = node.next
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.head == nil {
		if index == 0 {
			this.head = &Node{
				next: nil,
				val:  val,
			}
			return
		}
		return
	}

	node := this.head
	var i int

	// Iterate to (index - 1) node, if possible
	for i = 0; i < index-1; i++ {
		if node.next != nil {
			node = node.next
		} else {
			return
		}
	}

	node.next = &Node{
		next: node.next,
		val:  val,
	}

	return
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.head == nil {
		return
	}

	node := this.head
	var i int

	// Iterate to (index - 1) node, if possible
	for i = 0; i < index-1; i++ {
		if node.next != nil {
			node = node.next
		} else {
			return
		}
	}

	if node.next == nil {
		return
	}

	node.next = node.next.next
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
