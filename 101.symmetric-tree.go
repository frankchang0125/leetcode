/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 *
 * https://leetcode.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (43.37%)
 * Likes:    2172
 * Dislikes: 45
 * Total Accepted:    396.5K
 * Total Submissions: 914.1K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * Given a binary tree, check whether it is a mirror of itself (ie, symmetric
 * around its center).
 * 
 * For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
 * 
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 * 
 * 
 * 
 * 
 * But the following [1,2,2,null,3,null,3] is not:
 * 
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 * 
 * 
 * 
 * 
 * Note:
 * Bonus points if you could solve it both recursively and iteratively.
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
// Reference: http://bit.ly/2FA1ooE
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return symmetric(root.Left, root.Right)
}

func symmetric(node1 *TreeNode, node2 *TreeNode) bool {
    if node1 == nil && node2 == nil {
        return true
    }

    if node1 == nil {
        return false
    }

    if node2 == nil {
        return false
    }

    if node1.Val != node2.Val {
        return false
    }

    return symmetric(node1.Left, node2.Right) && symmetric(node1.Right, node2.Left)
}
