/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 *
 * https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (51.27%)
 * Likes:    1112
 * Dislikes: 39
 * Total Accepted:    221.1K
 * Total Submissions: 430.9K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * Given a binary search tree, write a function kthSmallest to find the kth
 * smallest element in it.
 * 
 * Note: 
 * You may assume k is always valid, 1 ≤ k ≤ BST's total elements.
 * 
 * Example 1:
 * 
 * 
 * Input: root = [3,1,4,null,2], k = 1
 * ⁠  3
 * ⁠ / \
 * ⁠1   4
 * ⁠ \
 * 2
 * Output: 1
 * 
 * Example 2:
 * 
 * 
 * Input: root = [5,3,6,2,4,null,null,1], k = 3
 * ⁠      5
 * ⁠     / \
 * ⁠    3   6
 * ⁠   / \
 * ⁠  2   4
 * ⁠ /
 * ⁠1
 * Output: 3
 * 
 * 
 * Follow up:
 * What if the BST is modified (insert/delete operations) often and you need to
 * find the kth smallest frequently? How would you optimize the kthSmallest
 * routine?
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
// Iteration version use stack to simulate
// the behavior of recursion version.
func kthSmallest(root *TreeNode, k int) int {
   if root == nil {
      return 0
   }

   stack := []*TreeNode{}
   node := root
   count := 0

   for node != nil || len(stack) > 0 {
      for node != nil {
         stack = append(stack, node) // Push
         node = node.Left
      }

      node = stack[len(stack) - 1]
      stack = stack[:len(stack) - 1] // Pop
      count++
      if count == k {
         return node.Val
      }

      node = node.Right
   }

   return 0
}
