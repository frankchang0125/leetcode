/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 *
 * https://leetcode.com/problems/binary-search/description/
 *
 * algorithms
 * Easy (47.39%)
 * Likes:    235
 * Dislikes: 30
 * Total Accepted:    48K
 * Total Submissions: 101.3K
 * Testcase Example:  '[-1,0,3,5,9,12]\n9'
 *
 * Given a sorted (in ascending order) integer array nums of n elements and a
 * target value, write a function to search target in nums. If target exists,
 * then return its index, otherwise return -1.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [-1,0,3,5,9,12], target = 9
 * Output: 4
 * Explanation: 9 exists in nums and its index is 4
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [-1,0,3,5,9,12], target = 2
 * Output: -1
 * Explanation: 2 does not exist in nums so return -1
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * You may assume that all elements in nums are unique.
 * n will be in the range [1, 10000].
 * The value of each element in nums will be in the range [-9999, 9999].
 * 
 * 
 */
func search(nums []int, target int) int {
    N := len(nums)
    if N == 0 {
        return -1
    }

    l := 0
    r := N - 1

    for l <= r {
        m := l + (r - l) / 2

        if nums[m] == target {
            return m
        } else if nums[m] > target {
            // Search left part.
            r = m - 1
        } else {
            // Search right part.
            l = m + 1
        }
    }

    return -1
}
