/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 *
 * https://leetcode.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (40.05%)
 * Total Accepted:    205.4K
 * Total Submissions: 508.4K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * Given an unsorted array of integers, find the length of longest increasing
 * subsequence.
 *
 * Example:
 *
 *
 * Input: [10,9,2,5,3,7,101,18]
 * Output: 4
 * Explanation: The longest increasing subsequence is [2,3,7,101], therefore
 * the length is 4.
 *
 * Note:
 *
 *
 * There may be more than one LIS combination, it is only necessary for you to
 * return the length.
 * Your algorithm should run in O(n^2) complexity.
 *
 *
 * Follow up: Could you improve it to O(n log n) time complexity?
 *
 */
// References:
//	http://bit.ly/2FzXcVN
//	http://bit.ly/2FzXcoL
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	maxLen := 1

	// For j = 0->i-1:
	//	If we can find a subsequence in range: nums[0->i-1], which nums[j] < nums[i],
	//	then nums[i] can be appended to such subsequence to form a
	//	longer increasing subsequence.
	//	=> dp[i] = The longest length of increasing subsequence in range: nums[0->i]
	//	=> dp[i] = max(dp[j]+1, dp[i]), where j = 0->i-1.
	//	=> dp[0] = 1
	//
	// ex: nums = [10, 9, 2, 5, 3, 7, 101, 18]
	//	=> i = 0, nums[0] = 10:  maxLen = 1: [10]
	//	=> i = 1, nums[1] = 9:   maxLen = 1: [10] or [9]
	//	=> i = 2, nums[2] = 2:   maxLen = 1: [10] or [9] or [2]
	//	=> i = 3, nums[3] = 5:   maxLen = 2: [2, 5]
	//	=> i = 4, nums[4] = 3:   maxLen = 2: [2, 5] or [2, 3]
	//	=> i = 5, nums[5] = 7:   maxLen = 3: [2, 5, 7] or [2, 3, 7]
	//	=> i = 6, nums[6] = 101: maxLen = 4: [2, 4, 7, 101] or [2, 3, 7, 101]
	//	=> i = 7, nums[7] = 18:  maxLen = 4: [2, 4, 7, 18] or [2, 3, 7, 18]
	for i := 1; i < len(nums); i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}

		maxLen = max(dp[i], maxLen)
	}

	return maxLen
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
