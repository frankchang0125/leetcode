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

	mem := make([]int, len(nums))

	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = max(lis(nums, i, mem), ans)
	}
	return ans
}

func lis(nums []int, i int, mem []int) int {
	if mem[i] != 0 {
		return mem[i]
	}

	if i == 0 {
		return 1
	}

	ans := 1
	for j := 0; j < i; j++ {
		if nums[j] < nums[i] {
			ans = max(lis(nums, j, mem) + 1, ans)
		}
	}

	mem[i] = ans
	return ans
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
