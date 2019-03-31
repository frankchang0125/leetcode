/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 *
 * https://leetcode.com/problems/triangle/description/
 *
 * algorithms
 * Medium (38.03%)
 * Total Accepted:    173.6K
 * Total Submissions: 448.9K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * Given a triangle, find the minimum path sum from top to bottom. Each step
 * you may move to adjacent numbers on the row below.
 *
 * For example, given the following triangle
 *
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
 *
 * Note:
 *
 * Bonus point if you are able to do this using only O(n) extra space, where n
 * is the total number of rows in the triangle.
 *
 */
// Triangle:
// [                 [
// ⁠    [2],           [2],
// ⁠   [3,4],          [3, 4],
// ⁠  [6,5,7],   ->    [6, 5, 7],
// ⁠ [4,1,8,3]         [4, 1, 8, 3]
// ]                 ]
//
// Reference: https://goo.gl/ZWNY2e
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 0 {
		return 0
	}

	// For each iteration, we only calculate
	// current level and previous level.
	// Thus, we can use rolling array technique to keep
	// the minimum path sums of previous level and current level.
	// Use two-level array:
	// 	dp[0]: Minimum path sums of previous level.
	//	dp[1]: Minimum path sums of current level.
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, m)
	}

	dp[0][0] = triangle[0][0]

	for y := 1; y < m; y++ {
		for x := 0; x < y+1; x++ {
			if x == 0 {
				// Left-most node,
				// must go from previous level's left-most node.
				dp[1][x] = dp[0][x] + triangle[y][x]
			} else if x == y {
				// Right-most node,
				// must go from previous level's right-most node.
				dp[1][x] = dp[0][x-1] + triangle[y][x]
			} else {
				// Choose the minimum path sum from previous level's
				// adjacent nodes.
				dp[1][x] = min(dp[0][x-1], dp[0][x]) + triangle[y][x]
			}
		}

		// Swap current level for next iteration.
		dp[0], dp[1] = dp[1], dp[0]
	}

	// Find minimum value of the bottom-most level.
	minVal := math.MaxInt32

	for i := 0; i < m; i++ {
		// We've swapped the last calculated minimum path sums
		// of bottom-most level to dp[0].
		minVal = min(dp[0][i], minVal)
	}

	return minVal
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
