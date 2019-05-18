/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 *
 * https://leetcode.com/problems/implement-stack-using-queues/description/
 *
 * algorithms
 * Easy (37.60%)
 * Likes:    317
 * Dislikes: 414
 * Total Accepted:    128K
 * Total Submissions: 327.4K
 * Testcase Example:  '["MyStack","push","push","top","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a stack using queues.
 * 
 * 
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * empty() -- Return whether the stack is empty.
 * 
 * 
 * Example:
 * 
 * 
 * MyStack stack = new MyStack();
 * 
 * stack.push(1);
 * stack.push(2);  
 * stack.top();   // returns 2
 * stack.pop();   // returns 2
 * stack.empty(); // returns false
 * 
 * Notes:
 * 
 * 
 * You must use only standard operations of a queue -- which means only push to
 * back, peek/pop from front, size, and is empty operations are valid.
 * Depending on your language, queue may not be supported natively. You may
 * simulate a queue by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a queue.
 * You may assume that all operations are valid (for example, no pop or top
 * operations will be called on an empty stack).
 * 
 * 
 */
type MyStack struct {
    q1  []int
    q2  []int
    top int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{
        q1: []int{},
        q2: []int{},
    }
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
    if len(this.q1) != 0 {
        this.q1 = append(this.q1, x)
    } else {
        this.q2 = append(this.q2, x)
    }
    this.top = x
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    var src *[]int
    var dst *[]int

    if len(this.q1) != 0 {
        src = &this.q1
        dst = &this.q2
    } else {
        src = &this.q2
        dst = &this.q1
    }

    for i := 0; i < len(*src)-1; i++ {
        *dst = append(*dst, (*src)[i])
    }

    if len(*dst) > 0 {
        this.top = (*dst)[len(*dst)-1]
    }

    pop := (*src)[len(*src)-1]
    *src = (*src)[:0]

    return pop
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.top
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    if len(this.q1) == 0 && len(this.q2) == 0 {
        return true
    }
    return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
