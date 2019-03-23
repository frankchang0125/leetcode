/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 *
 * https://leetcode.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (46.10%)
 * Total Accepted:    314.9K
 * Total Submissions: 667.7K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given a set of candidate numbers (candidates) (without duplicates) and a
 * target number (target), find all unique combinations in candidates where the
 * candidate numbers sums to target.
 *
 * The same repeated number may be chosen from candidates unlimited number of
 * times.
 *
 * Note:
 *
 *
 * All numbers (including target) will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [2,3,6,7], target = 7,
 * A solution set is:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,3,5], target = 8,
 * A solution set is:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 */
// Reference: https://goo.gl/bX7xKv
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	sort.Sort(sort.IntSlice(candidates))
	ans := [][]int{}
	dfs(candidates, target, 0, []int{}, 0, &ans)
	return ans
}

// s: Start index of candidates to be added to the solution.
// solution: Current combination of solution.
func dfs(candidates []int, target int, s int, solution []int, sum int, ans *[][]int) {
	if sum == target {
		s := make([]int, len(solution))
		copy(s, solution)
		*ans = append(*ans, s)
		return
	}

	for i := s; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			// Stop further iterations if sum+candidates[i]
			// has already exceeded the value of target.
			// (We've sorted candidates already.)
			return
		}

		solution = append(solution, candidates[i]) // Push
		// Different from pure combination problem,
		// we can use current candidate repeatedly,
		// thus pass 'i' instead of 'i+1' to next DFS.
		dfs(candidates, target, i, solution, sum+candidates[i], ans)
		solution = solution[:len(solution)-1] // Pop
	}
}
