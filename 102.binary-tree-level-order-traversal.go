/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (48.33%)
 * Likes:    1439
 * Dislikes: 36
 * Total Accepted:    372.8K
 * Total Submissions: 771.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
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
 * return its level order traversal as:
 * 
 * [
 * ⁠ [3],
 * ⁠ [9,20],
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
// Reference: http://bit.ly/2FA2pNl
func levelOrder(root *TreeNode) [][]int {
    ans := [][]int{}

    if root == nil {
        return ans
    }

    queue := []*TreeNode{root} // Push

    for len(queue) > 0 {
        level := []int{}

        for i := len(queue); i > 0; i-- {
            node := queue[0]
            queue = queue[1:] // Pop

            level = append(level, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left) // Push
            }

            if node.Right != nil {
                queue = append(queue, node.Right) // Push
            }
        }

        ans = append(ans, level)
    }

    return ans
}
