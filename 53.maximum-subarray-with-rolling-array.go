/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (42.51%)
 * Total Accepted:    489.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */
// For each number, nums[i]:
//  We can choose whether to pick previous sub-array or not.
// Therefore, in order to get the maximum sum of nums[0->i] sub-arrays:
// Don't pick nums[0->i-1] sub-array with negative 'sum'.
//
// Reference: https://goo.gl/h6Kr6R
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Use one number to record the maximum sum
	// of nums[0->i-1] sub-arrays.
	// Use another number to record the overall maximum sum
	// of sub-arrays.
	sum := nums[0]
	maxVal := sum
	for i := 1; i < len(nums); i++ {
		// If sums[i-1] < 0, nums[i] + sums[i-1] must < nums[i],
		// don't pick the sub-array.
		sum = max(nums[i]+sum, nums[i])
		maxVal = max(sum, maxVal)
	}

	return maxVal
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
