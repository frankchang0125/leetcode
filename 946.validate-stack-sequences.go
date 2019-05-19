/*
 * @lc app=leetcode id=946 lang=golang
 *
 * [946] Validate Stack Sequences
 *
 * https://leetcode.com/problems/validate-stack-sequences/description/
 *
 * algorithms
 * Medium (57.76%)
 * Likes:    232
 * Dislikes: 6
 * Total Accepted:    13K
 * Total Submissions: 22.5K
 * Testcase Example:  '[1,2,3,4,5]\n[4,5,3,2,1]'
 *
 * Given two sequences pushed and popped with distinct values, return true if
 * and only if this could have been the result of a sequence of push and pop
 * operations on an initially empty stack.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
 * Output: true
 * Explanation: We might do the following sequence:
 * push(1), push(2), push(3), push(4), pop() -> 4,
 * push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
 * Output: false
 * Explanation: 1 cannot be popped before 2.
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 0 <= pushed.length == popped.length <= 1000
 * 0 <= pushed[i], popped[i] < 1000
 * pushed is a permutation of popped.
 * pushed and popped have distinct values.
 * 
 * 
 * 
 */
// Rules of valid stack sequences:
//  1. If stack top == popped[poppedIdx], pop the stack.
//  2. If stack top != popped[poppedIdx], push all the numbers in pushed sequence
//     until Rule#1 satisfy.
//  3. After all numbers in pushed sequence are pushed, but still can't find a valid
//     number to pop out, then these stack sequences are invalid.
//  4. If all numbers in pushed sequence are pushed, all numbers in popped sequence
//     are popped and stack is empty, then these are valid stack sequences.
func validateStackSequences(pushed []int, popped []int) bool {
    N := len(pushed)
    stack := []int{}

    pushedIdx := 0
    poppedIdx := 0

    for poppedIdx < N {
        // If stack top != popped sequence number, push all the numbers in pushed sequence
        // until stack top == popped sequence number.
        for isEmpty(stack) || top(stack) != popped[poppedIdx] {
            if pushedIdx == N {
                break
            }

            push(&stack, pushed[pushedIdx])
            pushedIdx++
        }

        if top(stack) != popped[poppedIdx] {
            // Stack is not empty and stack top not equals to the popped sequence number.
            // No further steps we can do, these are not valid stack sequences.
            break
        }

        // Stack is not empty and stack top equals to the popped sequence number,
        // pop the stack.
        pop(&stack)
        poppedIdx++
    }

    if isEmpty(stack) && poppedIdx == N {
        return true
    }

    return false
}

func isEmpty(stack []int) bool {
    if len(stack) == 0 {
        return true
    }
    return false
}

func top(stack []int) int {
    if isEmpty(stack) {
        return -1
    }
    return stack[len(stack)-1]
}

func pop(stack *[]int) {
    if isEmpty(*stack) {
        return
    }
    *stack = (*stack)[:len(*stack)-1]
}

func push(stack *[]int, val int) {
    *stack = append(*stack, val)
}
