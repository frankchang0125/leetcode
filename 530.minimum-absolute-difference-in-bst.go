/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
 *
 * https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (49.54%)
 * Likes:    503
 * Dislikes: 35
 * Total Accepted:    57.6K
 * Total Submissions: 114.8K
 * Testcase Example:  '[1,null,3,2]'
 *
 * Given a binary search tree with non-negative values, find the minimum
 * absolute difference between values of any two nodes.
 * 
 * Example:
 * 
 * 
 * Input:
 * 
 * ⁠  1
 * ⁠   \
 * ⁠    3
 * ⁠   /
 * ⁠  2
 * 
 * Output:
 * 1
 * 
 * Explanation:
 * The minimum absolute difference is 1, which is the difference between 2 and
 * 1 (or between 2 and 3).
 * 
 * 
 * 
 * 
 * Note: There are at least two nodes in this BST.
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    nodes := []int{}
    inOrder(root, &nodes)

    minVal := nodes[len(nodes)-1]

    for i := 1; i < len(nodes); i++ {
        minVal = min(nodes[i] - nodes[i-1], minVal)
    }

    return minVal
}

func inOrder(node *TreeNode, nodes *[]int) {
    if node != nil {
        inOrder(node.Left, nodes)
        *nodes = append(*nodes, node.Val)
        inOrder(node.Right, nodes)
    }
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
