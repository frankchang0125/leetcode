/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 *
 * https://leetcode.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (37.69%)
 * Likes:    912
 * Dislikes: 294
 * Total Accepted:    310.7K
 * Total Submissions: 824K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * Given a binary tree and a sum, determine if the tree has a root-to-leaf path
 * such that adding up all the values along the path equals the given sum.
 * 
 * Note: A leaf is a node with no children.
 * 
 * Example:
 * 
 * Given the below binary tree and sum = 22,
 * 
 * 
 * ⁠     5
 * ⁠    / \
 * ⁠   4   8
 * ⁠  /   / \
 * ⁠ 11  13  4
 * ⁠/  \      \
 * 7    2      1
 * 
 * 
 * return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
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
func hasPathSum(root *TreeNode, sum int) bool {
    return dfs(root, 0, sum)
}

func dfs(node *TreeNode, cur int, target int) bool {
    if node == nil {
        return false
    }

    cur += node.Val
    if node.Left == nil && node.Right == nil && cur == target {
        return true
    }

    return dfs(node.Left, cur, target) || dfs(node.Right, cur, target)
}
