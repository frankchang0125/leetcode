/*
 * @lc app=leetcode id=540 lang=golang
 *
 * [540] Single Element in a Sorted Array
 *
 * https://leetcode.com/problems/single-element-in-a-sorted-array/description/
 *
 * algorithms
 * Medium (56.82%)
 * Total Accepted:    46.5K
 * Total Submissions: 81.8K
 * Testcase Example:  '[1,1,2]'
 *
 *
 * Given a sorted array consisting of only integers where every element appears
 * twice except for one element which appears once. Find this single element
 * that appears only once.
 *
 *
 * Example 1:
 *
 * Input: [1,1,2,3,3,4,4,8,8]
 * Output: 2
 *
 *
 *
 * Example 2:
 *
 * Input: [3,3,7,7,10,11,11]
 * Output: 10
 *
 *
 *
 * Note:
 * Your solution should run in O(log n) time and O(1) space.
 *
 *
 */
// Reference: http://bit.ly/2VhNtwx
func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	index := len(nums) / 2

	if index%2 != 0 {
		if nums[index] == nums[index-1] {
			// Must on the right part.
			return singleNonDuplicate(nums[index+1 : len(nums)])
		} else if nums[index] == nums[index+1] {
			// Must on the left part.
			return singleNonDuplicate(nums[0:index])
		}
	} else {
		if nums[index] == nums[index-1] {
			// Must on the left part.
			return singleNonDuplicate(nums[0 : index+1])
		} else if nums[index] == nums[index+1] {
			// Must on the right part.
			return singleNonDuplicate(nums[index:len(nums)])
		}
	}

	// It's index itself!!
	return nums[index]
}
