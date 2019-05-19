/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (41.59%)
 * Likes:    930
 * Dislikes: 58
 * Total Accepted:    216.4K
 * Total Submissions: 520.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the zigzag level order traversal of its nodes'
 * values. (ie, from left to right, then right to left for the next level and
 * alternate between).
 * 
 * 
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 
 * return its zigzag level order traversal as:
 * 
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
 * ]
 * 
 * 
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// For zigzag level order traversal, instead of using queue to preserve next level child nodes;
// since zigzag is LIFO for next level child nodes traversal, we can use 'two stacks'.
func zigzagLevelOrder(root *TreeNode) [][]int {
    ans := [][]int{}

    if root == nil {
        return ans
    }

    rightStack := []*TreeNode{}
    leftStack := []*TreeNode{}

    rightStack = append(rightStack, root) // Push
    rightDirection := true

    var s1 *[]*TreeNode
    var s2 *[]*TreeNode

    for {
        if rightDirection {
            s1 = &rightStack
            s2 = &leftStack
        } else {
            s1 = &leftStack
            s2 = &rightStack
        }

        if len(*s1) == 0 {
            break
        }

        level := []int{}

        for i := len(*s1) - 1; i >= 0; i-- {
            node := (*s1)[i]
            level = append(level, node.Val)

            if rightDirection {
                if node.Left != nil {
                    *s2 = append(*s2, node.Left) // Push
                }
        
                if node.Right != nil {
                    *s2 = append(*s2, node.Right) // Push
                }
            } else {
                if node.Right != nil {
                    *s2 = append(*s2, node.Right) // Push
                }
        
                if node.Left != nil {
                    *s2 = append(*s2, node.Left) // Push
                }
            }
        }

        *s1 = (*s1)[:0] // Clear the stack.
        rightDirection = !rightDirection
        ans = append(ans, level)
    }

    return ans   
}
