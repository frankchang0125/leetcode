/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 *
 * https://leetcode.com/problems/partition-equal-subset-sum/description/
 *
 * algorithms
 * Medium (39.58%)
 * Total Accepted:    79.3K
 * Total Submissions: 197.6K
 * Testcase Example:  '[1,5,11,5]'
 *
 * Given a non-empty array containing only positive integers, find if the array
 * can be partitioned into two subsets such that the sum of elements in both
 * subsets is equal.
 *
 * Note:
 *
 *
 * Each of the array element will not exceed 100.
 * The array size will not exceed 200.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1, 5, 11, 5]
 *
 * Output: true
 *
 * Explanation: The array can be partitioned as [1, 5, 5] and [11].
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1, 2, 3, 5]
 *
 * Output: false
 *
 * Explanation: The array cannot be partitioned into equal sum subsets.
 *
 *
 *
 *
 */
// Reference: http://bit.ly/2Vr5JAC
func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

    // To be partition equally, the sum must be an odd number.
	if sum%2 != 0 {
		return false
	}

	// This problem can be changed to:
	//	Can we find a subset of nums which its sum equals to: (sum of nums) / 2?
	sum /= 2

	// dp[j]: Can we find a subset of nums which sum equals to 'j'?
	dp := make([]bool, sum+1)

	// We can find a subset of nums which sum equals to '0'
	// i.e. Empty set. (Don't pick any numbers in nums.)
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			// 1. If dp[j-nums[i]] is true, then we are able to find a subset
			//    which sum equals to 'j' after adding nums[i] into it.
			// 2. If we already found a subset of nums which sum equals to 'j',
			//    then no matter how many numbers we choose, dp[j] must be true.
			dp[j] = dp[j-nums[i]] || dp[j]
		}
	}

	return dp[sum]
}
