/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (52.83%)
 * Total Accepted:    343.3K
 * Total Submissions: 641.1K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a collection of distinct integers, return all possible permutations.
 *
 * Example:
 *
 *
 * Input: [1,2,3]
 * Output:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 *
 */
// For permutation, [1, 2] and [2, 1] are different.
// Find all n-length permutations
// of given array nums with DISTINCT numbers.
func permute(nums []int) [][]int {
	ans := [][]int{}
	used := make([]bool, len(nums))

	// Find all n-length permutations recursively.
	dfs(&ans, nums, used, len(nums), []int{})
	return ans
}

// used: The array which indicate whether the number has been used or not
//       in current permutation.
// n: Target length of permutation.
// cur: Current permutation.
func dfs(ans *[][]int, nums []int, used []bool, n int, cur []int) {
	if len(cur) == n {
		// cur may be changed by other dfs,
		// create a copy and add it to the answers.
		c := make([]int, len(cur))
		copy(c, cur)
		*ans = append(*ans, c)
		return
	}

	// For permutation, since [1, 2] and [2, 1] are different,
	// we have to iterate from first number to last number in nums array
	// to find all permutations.
	for i := 0; i < len(nums); i++ {
		// Skip if the number is already used in current permutation.
		if used[i] {
			continue
		}

		// Append number to current permutation.
		cur = append(cur, nums[i])

		// Mark current number as used.
		used[i] = true

		dfs(ans, nums, used, n, cur)

		// We have to pop the appended number in order to prevent
		// affecting other dfs.
		cur = cur[:len(cur)-1]

		// We have to reset current number as non-used in order to prevent
		// affecting other dfs.
		used[i] = false
	}
}
