/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
 *
 * https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
 *
 * algorithms
 * Easy (57.71%)
 * Total Accepted:    74.3K
 * Total Submissions: 127.5K
 * Testcase Example:  '[3,9,20,15,7]'
 *
 * Given a non-empty binary tree, return the average value of the nodes on each
 * level in the form of an array.
 * 
 * Example 1:
 * 
 * Input:
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * Output: [3, 14.5, 11]
 * Explanation:
 * The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on
 * level 2 is 11. Hence return [3, 14.5, 11].
 * 
 * 
 * 
 * Note:
 * 
 * The range of node's value is in the range of 32-bit signed integer.
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
// Reference: http://bit.ly/2UI1BiB
func averageOfLevels(root *TreeNode) []float64 {
    return bfs(root)
}

func bfs(node *TreeNode) (avgs []float64) {
    avgs = []float64{}
    queue := []*TreeNode{node} // Enqueue

    for {
        if len(queue) == 0 {
            break
        }

        sum := 0
        size := len(queue)

        for i := size; i > 0; i-- {
            n := queue[0]
            queue = queue[1:] // Dequeue

            sum += n.Val

            if n.Left != nil {
                queue = append(queue, n.Left)
            }

            if n.Right != nil {
                queue = append(queue, n.Right)
            }
        }

        avg := float64(sum) / float64(size)
        avgs = append(avgs, avg)
    }

    return avgs
}
