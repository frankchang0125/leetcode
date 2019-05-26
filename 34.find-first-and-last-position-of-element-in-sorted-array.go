/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (33.48%)
 * Likes:    1545
 * Dislikes: 85
 * Total Accepted:    299.5K
 * Total Submissions: 893K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 * 
 * Your algorithm's runtime complexity must be in the order of O(log n).
 * 
 * If the target is not found in the array, return [-1, -1].
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 * 
 */
func searchRange(nums []int, target int) []int {
    first := findFirstPostion(nums, target)
    last := findLastPostion(nums, target)
    return []int{first, last}
}

func findFirstPostion(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    l := 0
    r := len(nums) - 1

    for l <= r {
        m := l + (r - l) / 2

        if nums[m] == target {
            if m == 0 {
                // nums[m] is already the first number, must be the first target.
                return m
            } else if m > 0 {
                if nums[m-1] != target {
                    return m
                }
                // nums[m] is not the first target, search left part.
                r = m - 1
            }
        } else if nums[m] > target {
            // Search left part
            r = m - 1
        } else {
            // Search right part
            l = m + 1
        }
    }

    return -1
}

func findLastPostion(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    l := 0
    r := len(nums) - 1

    for l <= r {
        m := l + (r - l) / 2

        if nums[m] == target {
            if m == len(nums) - 1 {
                // nums[m] is already the last number, must be the last target.
                return m
            } else if m < len(nums) - 1 {
                if nums[m+1] != target {
                    return m
                }
                // nums[m] is not the last target, search right part.
                l = m + 1
            }
        } else if nums[m] > target {
            // Search left part
            r = m - 1
        } else {
            // Search right part
            l = m + 1
        }
    }

    return -1
}
