/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 *
 * https://leetcode.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (41.13%)
 * Total Accepted:    189.8K
 * Total Submissions: 457.3K
 * Testcase Example:  '[1,2,2]'
 *
 * Given a collection of integers that might contain duplicates, nums, return
 * all possible subsets (the power set).
 * 
 * Note: The solution set must not contain duplicate subsets.
 * 
 * Example:
 * 
 * 
 * Input: [1,2,2]
 * Output:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 * 
 * 
 */
func subsetsWithDup(nums []int) [][]int {
	ans := [][]int{}

	// We have to sort nums array first in order to
	// prevent creating duplicate subsets.
	sort.Sort(sort.IntSlice(nums))

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
		// Skip if:
		// 1. Since nums array has been sorted,
		//    for numbers other than first number on the same level of tree,
		//    if current number is equal to previous number,
		//    it means all the sub-combinations are all already created by previous number.
		//    (We've sorted numbers already.)
		//    ex: [1, 1, 2]:
		//                  [ ]
		//               /   |   \
		//              1    1    2
		//             / \   └─────── Second 1: i > s && nums[0] == nums[1], skip.
		//            1   2
		//           /
		//          2
		if i > s && nums[i] == nums[i-1] {
			continue
		}

		// Append number to current combination.
		cur = append(cur, nums[i])

		dfs(ans, nums, i+1, n, cur)

		// We have to pop the appended number in order to prevent
		// affecting other dfs.
		cur = cur[:len(cur)-1]
	}
}
