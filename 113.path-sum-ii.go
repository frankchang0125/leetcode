/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 *
 * https://leetcode.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (40.49%)
 * Likes:    865
 * Dislikes: 32
 * Total Accepted:    229.5K
 * Total Submissions: 566.3K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * Given a binary tree and a sum, find all root-to-leaf paths where each path's
 * sum equals the given sum.
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
 * ⁠/  \    / \
 * 7    2  5   1
 * 
 * 
 * Return:
 * 
 * 
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
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
func pathSum(root *TreeNode, sum int) [][]int {
    ans := [][]int{}
    dfs(root, []int{}, 0, sum, &ans)
    return ans
}

func dfs(node *TreeNode, cur []int, curSum int, target int, ans *[][]int) {
    if node == nil {
        return
    }

    cur = append(cur, node.Val)
    curSum += node.Val
    if node.Left == nil && node.Right == nil && curSum == target {
        result := make([]int, len(cur))
        copy(result, cur)
        *ans = append(*ans, result)
        return
    }

    dfs(node.Left, cur, curSum, target, ans)
    dfs(node.Right, cur, curSum, target, ans)
}
