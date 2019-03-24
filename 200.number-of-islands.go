/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 *
 * https://leetcode.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (39.85%)
 * Total Accepted:    316.7K
 * Total Submissions: 780K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * Given a 2d grid map of '1's (land) and '0's (water), count the number of
 * islands. An island is surrounded by water and is formed by connecting
 * adjacent lands horizontally or vertically. You may assume all four edges of
 * the grid are all surrounded by water.
 *
 * Example 1:
 *
 *
 * Input:
 * 11110
 * 11010
 * 11000
 * 00000
 *
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input:
 * 11000
 * 11000
 * 00100
 * 00011
 *
 * Output: 3
 *
 */
// Reference: https://goo.gl/sEMfrT
func numIslands(grid [][]byte) int {
	N := len(grid)
	if N == 0 {
		return 0
	}

	M := len(grid[0])
	ans := 0

	// Iterate each land.
	for y := 0; y < N; y++ {
		for x := 0; x < M; x++ {
			if grid[y][x] == '0' {
				// Visited already, do nothing.
				continue
			}
			dfs(grid, M, N, x, y)

			// All walkable lands are visited,
			// we can determine an island now.
			ans++
		}
	}

	return ans
}

// Use DFS to visit all neighbor lands and its
// neighbor lands' neighbor lands of current land.
func dfs(grid [][]byte, M int, N int, x int, y int) {
	if x < 0 || x >= M || y < 0 || y >= N {
		// Out of bound, do nothing.
		return
	}

	if grid[y][x] == '0' {
		// Visited already, do nothing.
		return
	}
	grid[y][x] = '0'

	dfs(grid, M, N, x-1, y) // Go left
	dfs(grid, M, N, x, y-1) // Go up
	dfs(grid, M, N, x+1, y) // Go right
	dfs(grid, M, N, x, y+1) // Go down
}
