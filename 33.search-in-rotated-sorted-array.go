/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (32.86%)
 * Likes:    2374
 * Dislikes: 308
 * Total Accepted:    412.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 * 
 * (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 * 
 * You are given a target value to search. If found in the array return its
 * index, otherwise return -1.
 * 
 * You may assume no duplicate exists in the array.
 * 
 * Your algorithm's runtime complexity must be in the order of O(log n).
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 * 
 */
func search(nums []int, target int) int {
    l := 0
    r := len(nums) - 1

    for l <= r {
        m := l + (r - l) / 2

        if nums[m] == target {
            return m
        }

        if nums[m] >= nums[l] {
            if target < nums[m] && target >= nums[l] {
                r = m - 1
            } else {
                l = m + 1
            }
        } else {
            if target > nums[m] && target < nums[l] {
                l = m + 1
            } else {
                r = m - 1
            }
        }
    }

    return -1
}
