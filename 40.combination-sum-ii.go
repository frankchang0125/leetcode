/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 *
 * https://leetcode.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (39.69%)
 * Total Accepted:    206.1K
 * Total Submissions: 508.7K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * Given a collection of candidate numbers (candidates) and a target number
 * (target), find all unique combinations in candidates where the candidate
 * numbers sums to target.
 *
 * Each number in candidates may only be used once in the combination.
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
 * Input: candidates = [10,1,2,7,6,1,5], target = 8,
 * A solution set is:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,5,2,1,2], target = 5,
 * A solution set is:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 */
// Reference: https://goo.gl/t7eaU7
func combinationSum2(candidates []int, target int) [][]int {
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

		if i > s && candidates[i] == candidates[i-1] {
			// Stop further iterations if current number
			// equals to previous number in same depth.
			// (i.e. Disallow same number in same depth.)
			// (We've sorted candidates already.)
			continue
		}

		solution = append(solution, candidates[i]) // Push
		dfs(candidates, target, i+1, solution, sum+candidates[i], ans)
		solution = solution[:len(solution)-1] // Pop
	}
}
