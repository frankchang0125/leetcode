/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 *
 * https://leetcode.com/problems/combination-sum-iv/description/
 *
 * algorithms
 * Medium (43.81%)
 * Total Accepted:    80.1K
 * Total Submissions: 183.4K
 * Testcase Example:  '[1,2,3]\n4'
 *
 * Given an integer array with all positive numbers and no duplicates, find the
 * number of possible combinations that add up to a positive integer target.
 *
 * Example:
 *
 *
 * nums = [1, 2, 3]
 * target = 4
 *
 * The possible combination ways are:
 * (1, 1, 1, 1)
 * (1, 1, 2)
 * (1, 2, 1)
 * (1, 3)
 * (2, 1, 1)
 * (2, 2)
 * (3, 1)
 *
 * Note that different sequences are counted as different combinations.
 *
 * Therefore the output is 7.
 *
 *
 *
 *
 * Follow up:
 * What if negative numbers are allowed in the given array?
 * How does it change the problem?
 * What limitation we need to add to the question to allow negative numbers?
 *
 * Credits:
 * Special thanks to @pbrother for adding this problem and creating all test
 * cases.
 *
 */
// Recursively find the number of different combinations for target 'n',
// ex, for nums = [1, 2, 3], target = 4:
//  comb(4) = comb(4-1) + comb(4-2) + comb(4-3)
//          = comb(3) + comb(2) + comb(1)
//  comb(3) = comb(3-1) + comb(3-2) + comb(3-3)
//          = comb(2) + comb(1) + comb(0)
//  comb(2) = comb(2-1) + comb(2-2)
//          = comb(1) + comb(0)
//  comb(1) = comb(1-1)
//          = comb(0)
//  comb(0) = 1 (Only 1 combination of target 0: {})
// Use memoization to reduce the duplicate recursive calls.
//
// Reference: https://goo.gl/gzCdNz
func combinationSum4(nums []int, target int) int {
	mem := map[int]int{}
	mem[0] = 1
	return combination(nums, target, mem)
}

func combination(nums []int, target int, mem map[int]int) int {
	if target < 0 {
		return 0
	}

	if m, ok := mem[target]; ok {
		return m
	}

	var ans int
	for _, num := range nums {
		ans += combination(nums, target-num, mem)
	}

	mem[target] = ans
	return ans
}
