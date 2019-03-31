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
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])

	// Memoization, -1: Not calculated before.
	mem := make([][]int, m)
	for i := 0; i < m; i++ {
		mem[i] = make([]int, n)
	}

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			mem[y][x] = -1
		}
	}

	return uPaths(obstacleGrid, m-1, n-1, mem)
}

func uPaths(obstacleGrid [][]int, y, x int, mem [][]int) int {
	if x < 0 || y < 0 {
		// Out-of-bound, return 0 paths.
		return 0
	}

	if mem[y][x] != -1 {
		// Already calculated before, return the memorized result.
		return mem[y][x]
	}

	if x == 0 && y == 0 {
		// Initial value.
		if obstacleGrid[0][0] == 0 {
			return 1
		} else {
			return 0
		}
	}

	if obstacleGrid[y][x] == 1 {
		// Obstacle, cannot go, return 0 paths.
		mem[y][x] = 0
		return mem[y][x]
	}

	mem[y][x] = uPaths(obstacleGrid, y-1, x, mem) +
		uPaths(obstacleGrid, y, x-1, mem)
	return mem[y][x]
}
