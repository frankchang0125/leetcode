/*
 * @lc app=leetcode id=563 lang=golang
 *
 * [563] Binary Tree Tilt
 *
 * https://leetcode.com/problems/binary-tree-tilt/description/
 *
 * algorithms
 * Easy (46.87%)
 * Likes:    321
 * Dislikes: 717
 * Total Accepted:    50.9K
 * Total Submissions: 108.5K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a binary tree, return the tilt of the whole tree.
 * 
 * The tilt of a tree node is defined as the absolute difference between the
 * sum of all left subtree node values and the sum of all right subtree node
 * values. Null node has tilt 0.
 * 
 * The tilt of the whole tree is defined as the sum of all nodes' tilt.
 * 
 * Example:
 * 
 * Input: 
 * ⁠        1
 * ⁠      /   \
 * ⁠     2     3
 * Output: 1
 * Explanation: 
 * Tilt of node 2 : 0
 * Tilt of node 3 : 0
 * Tilt of node 1 : |2-3| = 1
 * Tilt of binary tree : 0 + 0 + 1 = 1
 * 
 * 
 * 
 * Note:
 * 
 * The sum of node values in any subtree won't exceed the range of 32-bit
 * integer. 
 * All the tilt values won't exceed the range of 32-bit integer.
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
func findTilt(root *TreeNode) int {
    ans := 0
    preorder(root, &ans)
    return ans
}

func preorder(node *TreeNode, ans *int) int {
    if node != nil {
        leftSum := preorder(node.Left, ans)
        rightSum := preorder(node.Right, ans)
        *ans += abs(leftSum - rightSum)
        return leftSum + rightSum + node.Val
    }
    return 0
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
} 
