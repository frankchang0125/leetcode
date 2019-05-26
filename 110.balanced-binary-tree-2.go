/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 *
 * https://leetcode.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (40.90%)
 * Likes:    1220
 * Dislikes: 108
 * Total Accepted:    319.4K
 * Total Submissions: 779.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, determine if it is height-balanced.
 * 
 * For this problem, a height-balanced binary tree is defined as:
 * 
 * 
 * a binary tree in which the depth of the two subtrees of every node never
 * differ by more than 1.
 * 
 * 
 * Example 1:
 * 
 * Given the following tree [3,9,20,null,null,15,7]:
 * 
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * Return true.
 * 
 * Example 2:
 * 
 * Given the following tree [1,2,2,3,3,null,null,4,4]:
 * 
 * 
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 * 
 * 
 * Return false.
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
// Use post-order traversal, when we traverse the root node,
// the tree height of left subtree and right subtree are already computed.
// Thus, we won't compute the tree heights redundantly.
func isBalanced(root *TreeNode) bool {
    height := 0
    return postOrder(root, &height)
}

func postOrder(node *TreeNode, height *int) bool {
    if node == nil {
        return true
    }

    leftHeight := 0
    leftBalanced := postOrder(node.Left, &leftHeight)

    rightHeight := 0
    rightBalanced := postOrder(node.Right, &rightHeight)

    if !leftBalanced || !rightBalanced {
        return false
    }

    if abs(leftHeight - rightHeight) > 1 {
        return false
    }

    *height = 1 + max(leftHeight, rightHeight)
    return true
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
