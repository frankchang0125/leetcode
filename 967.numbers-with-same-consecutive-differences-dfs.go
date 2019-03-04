/*
 * @lc app=leetcode id=967 lang=golang
 *
 * [967] Numbers With Same Consecutive Differences
 *
 * https://leetcode.com/problems/numbers-with-same-consecutive-differences/description/
 *
 * algorithms
 * Medium (35.72%)
 * Total Accepted:    5.4K
 * Total Submissions: 15.2K
 * Testcase Example:  '3\n7'
 *
 * Return all non-negative integers of length N such that the absolute
 * difference between every two consecutive digits is K.
 *
 * Note that every number in the answer must not have leading zeros except for
 * the number 0 itself. For example, 01 has one leading zero and is invalid,
 * but 0 is valid.
 *
 * You may return the answer in any order.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: N = 3, K = 7
 * Output: [181,292,707,818,929]
 * Explanation: Note that 070 is not a valid number, because it has leading
 * zeroes.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: N = 2, K = 1
 * Output: [10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= N <= 9
 * 0 <= K <= 9
 *
 *
 */
// Reference: https://goo.gl/Ms4esK
func numsSameConsecDiff(N int, K int) []int {
	solution := []int{}
	var start int

	if N == 1 {
		// If N == 1, '0' is also a valid answer.
		start = 0
	} else {
		// If N != 1, number started with 0 is not
		// allowed, ex: 010, 025, 037.. etc.
		// Thus, set start from 1.
		start = 1
	}

	// Find all valid numbers with DFS.
	for i := start; i <= 9; i++ {
		dfs(&solution, i, N-1, K)
	}

	return solution
}

func dfs(solution *[]int, ans int, N int, K int) {
	if N == 0 {
		// Total N-digits has found,
		// saved to the solution.
		*solution = append(*solution, ans)
		return
	}

	// Check whether (units digit + K) and (units digit - K)
	// are valid digits or not.
	u := ans % 10
	if u+K <= 9 {
		dfs(solution, ans*10+u+K, N-1, K)
	}

	// If K is 0, u+K and u-K will have same result,
	// skip u-K if K is 0.
	if K != 0 {
		if u-K >= 0 {
			dfs(solution, ans*10+u-K, N-1, K)
		}
	}
}
