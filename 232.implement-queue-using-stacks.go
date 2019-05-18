/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 *
 * https://leetcode.com/problems/implement-queue-using-stacks/description/
 *
 * algorithms
 * Easy (41.56%)
 * Likes:    524
 * Dislikes: 107
 * Total Accepted:    147.6K
 * Total Submissions: 342.3K
 * Testcase Example:  '["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a queue using stacks.
 * 
 * 
 * push(x) -- Push element x to the back of queue.
 * pop() -- Removes the element from in front of queue.
 * peek() -- Get the front element.
 * empty() -- Return whether the queue is empty.
 * 
 * 
 * Example:
 * 
 * 
 * MyQueue queue = new MyQueue();
 * 
 * queue.push(1);
 * queue.push(2);  
 * queue.peek();  // returns 1
 * queue.pop();   // returns 1
 * queue.empty(); // returns false
 * 
 * Notes:
 * 
 * 
 * You must use only standard operations of a stack -- which means only push to
 * top, peek/pop from top, size, and is empty operations are valid.
 * Depending on your language, stack may not be supported natively. You may
 * simulate a stack by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a stack.
 * You may assume that all operations are valid (for example, no pop or peek
 * operations will be called on an empty queue).
 * 
 * 
 */
// Reference: http://bit.ly/2FDTuuv
type MyQueue struct {
    s1 []int
    s2 []int
    front int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{
        s1: []int{},
        s2: []int{},
    }
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
    if len(this.s1) == 0 {
        this.front = x
    }
    this.s1 = append(this.s1, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    if len(this.s2) == 0 {
        for i := len(this.s1); i > 0; i-- {
            this.s2 = append(this.s2, this.s1[i-1])
        }
        this.s1 = this.s1[:0]
    }

    pop := this.s2[len(this.s2)-1]
    this.s2 = this.s2[:len(this.s2)-1]

    if len(this.s2) > 0 {
        this.front = this.s2[len(this.s2)-1]
    }

    return pop
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.front
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    if len(this.s1) == 0 && len(this.s2) == 0 {
        return true
    }
    return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
