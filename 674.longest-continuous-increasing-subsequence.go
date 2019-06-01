/*
 * @lc app=leetcode id=674 lang=golang
 *
 * [674] Longest Continuous Increasing Subsequence
 *
 * https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/
 *
 * algorithms
 * Easy (43.62%)
 * Total Accepted:    61.6K
 * Total Submissions: 140.1K
 * Testcase Example:  '[1,3,5,4,7]'
 *
 *
 * Given an unsorted array of integers, find the length of longest continuous
 * increasing subsequence (subarray).
 *
 *
 * Example 1:
 *
 * Input: [1,3,5,4,7]
 * Output: 3
 * Explanation: The longest continuous increasing subsequence is [1,3,5], its
 * length is 3.
 * Even though [1,3,5,7] is also an increasing subsequence, it's not a
 * continuous one where 5 and 7 are separated by 4.
 *
 *
 *
 * Example 2:
 *
 * Input: [2,2,2,2,2]
 * Output: 1
 * Explanation: The longest continuous increasing subsequence is [2], its
 * length is 1.
 *
 *
 *
 * Note:
 * Length of the array will not exceed 10,000.
 *
 */
// l[i]: Length of continuous increasing sequence: nums[0->i]
//
// If current number is larger than previous number,
// current number can be added to the sequence, l[i] = l[i-1] + 1.
// 
// If current number is smaller or equal to previous number,
// new sequence created, l[i] = 1.
//
// Reference: https://goo.gl/PLn8BW
func findLengthOfLCIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    l := make([]int, len(nums))

	l[0] = 1
    maxLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
            // Current number is smaller or equal to previous number,
            // new sequence created.
			l[i] = 1
		} else {
            // Current number is larger than previous number,
            // current number can be added to the sequence.
			l[i] = l[i-1] + 1
        }
        
        // Record the longest length of continuous increasing subsequences.
        maxLen = max(l[i], maxLen)
	}

    return maxLen
}

func max(x, y int) int {
    if x >= y {
        return x
    }
    return y
}
