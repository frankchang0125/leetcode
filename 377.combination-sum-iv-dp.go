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
// Use dynamic programming to find the number of different combinations for target 'n',
// dp[i]: Number of combinations for target 'i'
// ex, for nums = [1, 2, 3], target = 4:
//  dp[0] = 1
//  dp[1] = dp[1-1]
//  dp[2] = dp[2-1] + dp[2-2]
//  dp[3] = dp[3-1] + dp[3-2] + dp[3-3]
//  dp[4] = dp[4-1] + dp[4-2] + dp[4-3]
//
// Reference: https://goo.gl/gzCdNz
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}
