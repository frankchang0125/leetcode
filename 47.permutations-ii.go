/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 *
 * https://leetcode.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (38.74%)
 * Total Accepted:    223.3K
 * Total Submissions: 568.8K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a collection of numbers that might contain duplicates, return all
 * possible unique permutations.
 *
 * Example:
 *
 *
 * Input: [1,1,2]
 * Output:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 *
 */
func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	used := make([]bool, len(nums))

	// We have to sort nums array first in order to
	// prevent creating duplicate permutations.
	sort.Sort(sort.IntSlice(nums))

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
		// Skip if:
		// 1. The number is already used in current permutation.
		// 2. Since nums array has been sorted,
		//    if previous number was not used (i.e. two numbers are at same level)
		//	  but current number is equal to previous number,
		//    it means all the sub-permutations are already created by previous number.
		//    (We've sorted numbers already.)
		//    ex: [1, 1, 2]:
		//                  [ ]
		//               /   |   \
		//              1    1    2
		//             / \   └─────── Second 1: used[0] = false && nums[0] == nums[1], skip.
		//            1   2
		//           /     \
		//          2       1
		if used[i] || (i > 0 && !used[i-1] && nums[i] == nums[i-1]) {
			continue
		}

		// Mark current number as used.
		used[i] = true
		cur = append(cur, nums[i])
		dfs(ans, nums, used, n, cur)

		// We have to pop the appended number in order to prevent
		// affecting other dfs.
		cur = cur[:len(cur)-1]

		// We have to reset current number as non-used in order to prevent
		// affecting other dfs.
		used[i] = false
	}
}
