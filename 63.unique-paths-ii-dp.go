/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 *
 * https://leetcode.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (33.22%)
 * Total Accepted:    188.5K
 * Total Submissions: 566.8K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 *
 * Now consider if some obstacles are added to the grids. How many unique paths
 * would there be?
 *
 *
 *
 * An obstacle and empty space is marked as 1 and 0 respectively in the grid.
 *
 * Note: m and n will be at most 100.
 *
 * Example 1:
 *
 *
 * Input:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * Output: 2
 * Explanation:
 * There is one obstacle in the middle of the 3x3 grid above.
 * There are two ways to reach the bottom-right corner:
 * 1. Right -> Right -> Down -> Down
 * 2. Down -> Down -> Right -> Right
 *
 *
 */
// Reference: https://goo.gl/Q2F6fU
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	if n == 0 {
		return 0
	}

	m := len(obstacleGrid[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

    if obstacleGrid[0][0] == 1 {
		// Obstacle, cannot go, return 0 paths.
        return 0
    }

	// Initial value.
	dp[0][0] = 1

	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			if x == 0 && y == 0 {
				// Skip (0, 0) case, which is the initial value.
				continue
			}

			if obstacleGrid[y][x] == 1 {
				dp[y][x] = 0
				continue
			}

			if x == 0 {
				// Can only go from up to down.
				dp[y][x] = dp[y-1][x]
			} else if y == 0 {
				// Can only go from left to right.
				dp[y][x] = dp[y][x-1]
			} else {
				// Can go from up to down and left to right.
				dp[y][x] = dp[y-1][x] + dp[y][x-1]
			}
		}
	}

	return dp[n-1][m-1]
}
