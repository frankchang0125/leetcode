/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 *
 * https://leetcode.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (45.97%)
 * Total Accepted:    268.1K
 * Total Submissions: 573.4K
 * Testcase Example:  '3\n2'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 *
 * How many possible unique paths are there?
 *
 *
 * Above is a 7 x 3 grid. How many possible unique paths are there?
 *
 * Note: m and n will be at most 100.
 *
 * Example 1:
 *
 *
 * Input: m = 3, n = 2
 * Output: 3
 * Explanation:
 * From the top-left corner, there are a total of 3 ways to reach the
 * bottom-right corner:
 * 1. Right -> Right -> Down
 * 2. Right -> Down -> Right
 * 3. Down -> Right -> Right
 *
 *
 * Example 2:
 *
 *
 * Input: m = 7, n = 3
 * Output: 28
 *
 */
// Reference: https://goo.gl/4kk5k1
func uniquePaths(m int, n int) int {
	mem := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		mem[i] = make([]int, n+1)
	}

	return uPaths(m, n, mem)
}

func uPaths(m int, n int, mem [][]int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	if m == 1 && n == 1 {
		return 1
	}

	if mem[m][n] != 0 {
		return mem[m][n]
	}

    mem[m][n] = uPaths(m-1, n, mem) + uPaths(m, n-1, mem)
    return mem[m][n]
}
