/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 *
 * https://leetcode.com/problems/subsets/description/
 *
 * algorithms
 * Medium (50.36%)
 * Total Accepted:    331K
 * Total Submissions: 649.4K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a set of distinct integers, nums, return all possible subsets (the
 * power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: nums = [1,2,3]
 * Output:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */
// Reference: https://goo.gl/vtrHjm
// Find all subsets of DISTINCT numbers in nums array.
func subsets(nums []int) [][]int {
	ans := [][]int{}

	// Equals to find C(m, n), where n:
	// 	0 (Empty subset) -> m (Length of nums, subset contains all numbers in nums).
	// Find all different-length subsets recursively.
	for i := 0; i <= len(nums); i++ {
		dfs(&ans, nums, 0, i, []int{})
	}

	return ans
}

// s: Start index of nums to be added to the combination.
// n: Target length of combination.
// cur: Current combination.
func dfs(ans *[][]int, nums []int, s int, n int, cur []int) {
	if len(cur) == n {
		// cur may be changed by other dfs,
		// create a copy and add it to the answers.
		c := make([]int, len(cur))
		copy(c, cur)
		*ans = append(*ans, c)
		return
	}

	// Add number to the combination and recursively call DFS
	// to find all combinations.
	for i := s; i < len(nums); i++ {
		cur = append(cur, nums[i])
		dfs(ans, nums, i+1, n, cur)
		// We have to pop the appended number in order to prevent
		// affecting other dfs.
		cur = cur[:len(cur)-1]
	}
}
