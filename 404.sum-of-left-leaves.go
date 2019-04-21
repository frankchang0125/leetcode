/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
 *
 * https://leetcode.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (48.59%)
 * Total Accepted:    120.2K
 * Total Submissions: 246.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Find the sum of all left leaves in a given binary tree.
 *
 * Example:
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * There are two left leaves in the binary tree, with values 9 and 15
 * respectively. Return 24.
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
// Reference: http://bit.ly/2UGloio
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	PreOrder(root, &sum)
	return sum
}

func PreOrder(node *TreeNode, sum *int) {
	if node != nil {
        // If current node has left node and it's left node has no child nodes,
        // then current node's left node is a left leaf node.
		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				*sum += node.Left.Val
			}
		}
		PreOrder(node.Left, sum)
		PreOrder(node.Right, sum)
	}
}
