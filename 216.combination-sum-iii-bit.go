/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 *
 * https://leetcode.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (50.12%)
 * Total Accepted:    115.2K
 * Total Submissions: 227.2K
 * Testcase Example:  '3\n7'
 *
 *
 * Find all possible combinations of k numbers that add up to a number n, given
 * that only numbers from 1 to 9 can be used and each combination should be a
 * unique set of numbers.
 *
 * Note:
 *
 *
 * All numbers will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: k = 3, n = 7
 * Output: [[1,2,4]]
 *
 *
 * Example 2:
 *
 *
 * Input: k = 3, n = 9
 * Output: [[1,2,6], [1,3,5], [2,3,4]]
 *
 *
 */
// Reference: https://goo.gl/D29zXP
func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}

	// Use 9-bit binary to represent the combination.
	// Generate all combinations of 1->9 sequence.
	// (Total 2^9 combinations.)
	for i := 0; i < (1 << 9); i++ {
		sum := 0
		solution := []int{}

		for j := 1; j <= 9; j++ {
			if i&(1<<(uint(j)-1)) > 0 {
				solution = append(solution, j)
				sum += j
			}
		}

		if len(solution) == k && sum == n {
			ans = append(ans, solution)
		}
	}

	return ans
}
