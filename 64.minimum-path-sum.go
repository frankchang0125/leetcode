/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 *
 * https://leetcode.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (45.12%)
 * Total Accepted:    218.1K
 * Total Submissions: 473.1K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top
 * left to bottom right which minimizes the sum of all numbers along its path.
 *
 * Note: You can only move either down or right at any point in time.
 *
 * Example:
 *
 *
 * Input:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * Output: 7
 * Explanation: Because the path 1→3→1→1→1 minimizes the sum.
 *
 *
 */
// Reference: http://bit.ly/2UFWj7e
func minPathSum(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}

	m := len(grid[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	dp[0][0] = grid[0][0]

	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			if x == 0 && y == 0 {
				// Skip (0, 0) case, which is the initial value.
				continue
			} else if y == 0 {
				dp[y][x] = dp[y][x-1]
			} else if x == 0 {
				dp[y][x] = dp[y-1][x]
			} else {
				dp[y][x] = min(dp[y][x-1], dp[y-1][x])
			}

			dp[y][x] += grid[y][x]
		}
	}

	return dp[n-1][m-1]
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
