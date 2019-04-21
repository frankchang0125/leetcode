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

	// dp[s][y][x]: Number of paths to move the ball 's' times
	//              out of grid boundary from point: (y, x).
	// dp[0][y][x]: Move ball 0 times, number of paths is zero.
	//
	// If y < 0 || y >= m || x < 0 || x >= n: out of grid boundary,
    // number of paths = 1.
    //
    // dp[s][y][x] = dp[s-1][y][x-1] + (Go from Left)
    //               dp[s-1][y+1][x] + (Go from Down)
    //               dp[s-1][y][x+1] + (Go from Right)
    //               dp[s-1][y-1][x]   (Go from Up)
	dp := make([][][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, n)
		}
	}

	// directions of (y, x)
	directions := [][]int{
		[]int{0, -1}, // Go from Left
		[]int{1, 0},  // Go from Down
		[]int{0, 1},  // Go from Right
		[]int{-1, 0}, // Go from Up
	}

	for s := 1; s <= N; s++ {
		for y := 0; y < m; y++ {
			for x := 0; x < n; x++ {
				for _, d := range directions {
					newY := y + d[0]
					newX := x + d[1]

					if newY < 0 || newY >= m || newX < 0 || newX >= n {
						dp[s][y][x] += 1
					} else {
						dp[s][y][x] += dp[s-1][newY][newX]
						dp[s][y][x] %= mod
					}
				}
			}
		}
	}

	return dp[N][i][j]
}
