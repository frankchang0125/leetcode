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
	dfs(k, n, 1, []int{}, 0, &ans)
	return ans
}

// s: Start number of 1->9 sequence to be added to the solution.
// solution: Current combination of solution.
func dfs(k int, n int, s int, solution []int, sum int, ans *[][]int) {
	if k == 0 && sum == n {
		c := make([]int, len(solution))
		copy(c, solution)
		*ans = append(*ans, c)
		return
	}

	for i := s; i <= 9; i++ {
		if sum+i > n {
			// Stop further iterations if sum+i
			// has already exceeded the value of target.
			return
		}

		solution = append(solution, i) // Push
		dfs(k-1, n, i+1, solution, sum+i, ans)
		solution = solution[:len(solution)-1] // Pop
	}
}
