/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 *
 * https://leetcode.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (40.70%)
 * Total Accepted:    303K
 * Total Submissions: 741.4K
 * Testcase Example:  '[1,2,3,1]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed, the only constraint stopping
 * you from robbing each of them is that adjacent houses have security system
 * connected and it will automatically contact the police if two adjacent
 * houses were broken into on the same night.
 *
 * Given a list of non-negative integers representing the amount of money of
 * each house, determine the maximum amount of money you can rob tonight
 * without alerting the police.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money =
 * 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 * Example 2:
 *
 *
 * Input: [2,7,9,3,1]
 * Output: 12
 * Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house
 * 5 (money = 1).
 * Total amount you can rob = 2 + 9 + 1 = 12.
 *
 *
 */
// Reference: http://bit.ly/2VdtkYx
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp1 := nums[0]
	dp2 := 0

	for i := 1; i < len(nums); i++ {
		if i == 1 {
			dp2 = max(nums[i], dp1)
		} else {
			dp2 = max(dp2+nums[i], dp1)
		}
		dp1, dp2 = dp2, dp1
	}

	return dp1
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
