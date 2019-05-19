/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (48.28%)
 * Likes:    869
 * Dislikes: 41
 * Total Accepted:    257.2K
 * Total Submissions: 532.6K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the postorder traversal of its nodes' values.
 * 
 * Example:
 * 
 * 
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 * 
 * Output: [3,2,1]
 * 
 * 
 * Follow up: Recursive solution is trivial, could you do it iteratively?
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
// Reference: http://bit.ly/2Qam3E1
func postorderTraversal(root *TreeNode) []int {
    ans := []int{}
    postOrder(root, &ans)
    return ans
}

func postOrder(node *TreeNode, ans *[]int) {
    if node != nil {
        postOrder(node.Left, ans)
        postOrder(node.Right, ans)
        *ans = append(*ans, node.Val)
    }
}
