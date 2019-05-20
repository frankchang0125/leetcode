/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 *
 * https://leetcode.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (45.83%)
 * Likes:    853
 * Dislikes: 66
 * Total Accepted:    223.6K
 * Total Submissions: 487.8K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * Given a binary tree, return all root-to-leaf paths.
 * 
 * Note: A leaf is a node with no children.
 * 
 * Example:
 * 
 * 
 * Input:
 * 
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 * 
 * Output: ["1->2->5", "1->3"]
 * 
 * Explanation: All root-to-leaf paths are: 1->2->5, 1->3
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
func binaryTreePaths(root *TreeNode) []string {
    ans := []string{}
    dfs(root, "", &ans)
    return ans
}

func dfs(node *TreeNode, cur string, ans *[]string) {
    if node == nil {
        return
    }

    if cur == "" {
        // Root node.
        cur = fmt.Sprintf("%d", node.Val)
    } else {
        cur = fmt.Sprintf("%s->%d", cur, node.Val)
    }

    if node.Left == nil && node.Right == nil {
        *ans = append(*ans, cur)
        return
    }

    dfs(node.Left, cur, ans)
    dfs(node.Right, cur, ans)
}
