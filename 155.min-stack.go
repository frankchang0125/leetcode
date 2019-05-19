/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (36.85%)
 * Likes:    1731
 * Dislikes: 179
 * Total Accepted:    296.9K
 * Total Submissions: 805.6K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 * 
 * 
 * push(x) -- Push element x onto stack.
 * 
 * 
 * pop() -- Removes the element on top of the stack.
 * 
 * 
 * top() -- Get the top element.
 * 
 * 
 * getMin() -- Retrieve the minimum element in the stack.
 * 
 * 
 * 
 * 
 * Example:
 * 
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> Returns -3.
 * minStack.pop();
 * minStack.top();      --> Returns 0.
 * minStack.getMin();   --> Returns -2.
 * 
 * 
 */
type MinStack struct {
    s1 []int
    s2 []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        s1: []int{},
        s2: []int{},
    }
}


func (this *MinStack) Push(x int) {
    this.s1 = append(this.s1, x)

    if len(this.s2) == 0 || this.s2[len(this.s2)-1] >= x {
        this.s2 = append(this.s2, x)
    }
}


func (this *MinStack) Pop() {
    if len(this.s1) > 0 {
        pop := this.s1[len(this.s1)-1]
        this.s1 = this.s1[:len(this.s1)-1]
        if pop == this.s2[len(this.s2)-1] {
            this.s2 = this.s2[:len(this.s2)-1]
        }
    }
}


func (this *MinStack) Top() int {
    if len(this.s1) > 0 {
        return this.s1[len(this.s1)-1]
    }
    return -1
}


func (this *MinStack) GetMin() int {
    if len(this.s2) > 0 {
        return this.s2[len(this.s2)-1]
    }
    return -1
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
