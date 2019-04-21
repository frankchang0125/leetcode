/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 *
 * https://leetcode.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (40.35%)
 * Total Accepted:    385.2K
 * Total Submissions: 946.8K
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * Given a sorted array and a target value, return the index if the target is
 * found. If not, return the index where it would be if it were inserted in
 * order.
 *
 * You may assume no duplicates in the array.
 *
 * Example 1:
 *
 *
 * Input: [1,3,5,6], 5
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [1,3,5,6], 2
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: [1,3,5,6], 7
 * Output: 4
 *
 *
 * Example 4:
 *
 *
 * Input: [1,3,5,6], 0
 * Output: 0
 *
 *
 */
func searchInsert(nums []int, target int) int {
    return lowerBound(nums, target)
}

// Use binary search:
// 	lowerBound returns the first index: i, such that nums[i] >= target.
func lowerBound(nums []int, target int) int {
	// Left included / Right included.
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := l + (r-l)/2

		if nums[m] >= target {
			// Searched number is larger or equals to target,
			// search Left part.
			//  Since we want to find the lower bound of the target,
			//  thus, search Left part if searched number is larger
			//  or equals to target.
			r = m - 1
		} else {
			// Searched number is smaller than target,
			// search Right part.
			l = m + 1
		}
	}

	// 'l' is the first index, such that nums[l] >= target.
	return l
}
