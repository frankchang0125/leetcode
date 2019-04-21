/*
 * @lc app=leetcode id=576 lang=golang
 *
 * [576] Out of Boundary Paths
 *
 * https://leetcode.com/problems/out-of-boundary-paths/description/
 *
 * algorithms
 * Medium (31.09%)
 * Total Accepted:    16.2K
 * Total Submissions: 51.2K
 * Testcase Example:  '2\n2\n2\n0\n0'
 *
 * There is an m by n grid with a ball. Given the start coordinate (i,j) of the
 * ball, you can move the ball to adjacent cell or cross the grid boundary in
 * four directions (up, down, left, right). However, you can at most move N
 * times. Find out the number of paths to move the ball out of grid boundary.
 * The answer may be very large, return it after mod 10^9 + 7.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: m = 2, n = 2, N = 2, i = 0, j = 0
 * Output: 6
 * Explanation:
 *
 *
 *
 * Example 2:
 *
 *
 * Input: m = 1, n = 3, N = 3, i = 0, j = 1
 * Output: 12
 * Explanation:
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * Once you move the ball out of boundary, you cannot move it back.
 * The length and height of the grid is in range [1,50].
 * N is in range [0,50].
 *
 *
 */
// Reference: http://bit.ly/2Uz7IW7
func findPaths(m int, n int, N int, i int, j int) int {
	mod := int(math.Pow(10, 9)) + 7

	// dp[y][x]: Number of paths to move the ball 's' times
	//           out of grid boundary from point: (y, x).
	//
	// If y < 0 || y >= m || x < 0 || x >= n: out of grid boundary,
	// number of paths = 1.
	//
	// cur[y][x] = dp[y][x-1] + (Go from Left)
	//             dp[y+1][x] + (Go from Down)
	//             dp[y][x+1] + (Go from Right)
	//             dp[y-1][x]   (Go from Up)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// directions of (y, x)
	directions := [][]int{
		[]int{0, -1}, // Go from Left
		[]int{1, 0},  // Go from Down
		[]int{0, 1},  // Go from Right
		[]int{-1, 0}, // Go from Up
	}

	for s := 1; s <= N; s++ {
		cur := make([][]int, m)
		for i := 0; i < m; i++ {
			cur[i] = make([]int, n)
		}

		for y := 0; y < m; y++ {
			for x := 0; x < n; x++ {
				for _, d := range directions {
					newY := y + d[0]
					newX := x + d[1]

					if newY < 0 || newY >= m || newX < 0 || newX >= n {
						cur[y][x] += 1
					} else {
						cur[y][x] += dp[newY][newX]
						cur[y][x] %= mod
					}
				}
			}
		}

		cur, dp = dp, cur
	}

	return dp[i][j]
}
