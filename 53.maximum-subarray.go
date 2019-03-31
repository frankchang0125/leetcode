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
// sums[i]: Maximum sum of nums[0->i] sub-arrays, with nums[i] must be used.
// For each number, nums[i]:
//  We can choose whether to pick previous sub-array or not.
// Therefore, in order to get the maximum sum of nums[0->i] sub-arrays:
// Don't pick nums[0->i-1] sub-array with negative 'sum'.
//  ex:   [-2, 1, -3, 4, -1, 2, 1, -5, 4]
//  sums: [-2, 1, -2, 4,  3, 5, 6,  1, 5]
//  For sums[1]: nums[1] = 1, previous sum: sums[0] = -2, don't pick it.
//  For sums[2]: nums[2] = -3, previous sum: sums[1] = 1, pick it.
//  For sums[3]: nums[3] = 4, previous sum: sums[2] = -2, don't pick it.
//  ... etc
//
// Reference: https://goo.gl/h6Kr6R
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sums := make([]int, len(nums))

	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		// If sums[i-1] < 0, nums[i] + sums[i-1] must < nums[i],
		// don't pick the sub-array.
		sums[i] = max(nums[i]+sums[i-1], nums[i])
	}

    // Find the maximum sum of sub-arrays.
	maxVal := math.MinInt32

	for i := 0; i < len(sums); i++ {
		maxVal = max(sums[i], maxVal)
	}

	return maxVal
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
