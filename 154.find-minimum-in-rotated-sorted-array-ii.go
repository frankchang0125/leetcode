/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 *
 * https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Hard (38.83%)
 * Likes:    428
 * Dislikes: 138
 * Total Accepted:    129K
 * Total Submissions: 328.4K
 * Testcase Example:  '[1,3,5]'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 * 
 * (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
 * 
 * Find the minimum element.
 * 
 * The array may contain duplicates.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,3,5]
 * Output: 1
 * 
 * Example 2:
 * 
 * 
 * Input: [2,2,2,0,1]
 * Output: 0
 * 
 * Note:
 * 
 * 
 * This is a follow up problem to Find Minimum in Rotated Sorted Array.
 * Would allow duplicates affect the run-time complexity? How and why?
 * 
 * 
 */
// For rotated sorted array, we can divide array into left and right part,
// which the numbers in left and right part array are sorted individually.
//  ex: [3, 4, 5, 1, 2]:
//      => Left part:  [3, 4, 5]
//      => Right part: [1, 2]
//
//  ex: [5, 1, 2, 3, 4]:
//      => Left part:  [5]
//      => Right part: [1, 2, 3, 4]
//
// Therefore, we can use binary search, assume:
//  l: Index of left bound number.
//  r: Index of right bound number.
//  m: Index of middle number of left and right bound.
//
//  If m is in left part:
//      m will be equal or greater than l, since numbers in left part array are sorted.
//      The smallest number should be the number at right-hand side of m.
//          ex: [3, 4, 5, 1, 2]:
//              => Left part:  [3, 4, 5]
//              => Right part: [1, 2]
//              => Middle number: 5
//              => The smallest number should be the number at right-hand side of : [5, 1, 2] (m included)
//  If m is in right part:
//      m will be equal or smaller than r, since numbers in right part array are sorted.
//      The smallest number should be the number at left-hand side of m.
//          ex: [5, 1, 2, 3, 4]:
//              => Left part:  [5]
//              => Right part: [1, 2, 3, 4]
//              => Middle number: 2
//              => The smallest number should be the number at left-hand side of m: [5, 1, 2] (m included)
//  Special case: Middle number is the smallest number:
//      ex: [4, 5, 1, 2, 3]:
//          => Left part:  [4, 5]
//          => Right part: [1, 2, 3]
//          => Middle number: 1
//          => The smallest number can be the number at left-hand side of m: [4, 5, 1] (m included)
//          => The smallest number can also be the number at right-hand side of m: [1, 2, 3] (m included)
//
//  For array with duplicated values, we may not able to determine whether the middle number belongs to
//  left part array or right part array.
//       ex: [1, 0, 1, 1, 1] vs. [1, 1, 1, 0, 1]:
//          => nums[l], nums[r] and nums[m] are all equal to 1.
//          => For case 1: Middle number belongs to right part array [0, 1, 1, 1].
//          => However, for case 2: Middle number belongs to left part array [1, 1, 1].
//          => Under such case, we cannot determine whether the middle number belongs to
//             left part array or right part array. Thus we have no way but use iteration to find out
//             the minimum number.
//
// Try to move l and r until: nums[l] < nums[r], then m will be the index of the smallest number.
func findMin(nums []int) int {
    if len(nums) == 0 {
        return -1
    }

    if len(nums) == 1 {
        return nums[0]
    }

    l := 0
    r := len(nums) - 1
    // Special case, already sorted array: [1, 2, 3], smallest number = nums[l].
    m := l

    // m will eventually be the index of the smallest number.
    for nums[l] >= nums[r] {
        // Special case: only two numbers left, the smallest number must be nums[r].
        // Ex: [2, 1]
        // (P.S. [1, 2] won't enter for loop.)
        if r - l == 1 {
            m = r
            break
        }

        m = l + (r - l) / 2

        // nums[l], nums[r] and nums[m] are all equal.
        // We cannot determine whether the middle number belongs to
        // left part array or right part array.
        // Thus we have no way but use iteration to find out
        // the minimum number.
        if nums[l] == nums[r] && nums[m] == nums[l] {
            result := nums[l]
            for i := l + 1; i <= r; i++ {
                result = min(nums[i], result)
            }
            return result
        }

        if nums[m] >= nums[l] {
            // The smallest number should be the number at right-hand side of m.
            l = m
        } else if nums[m] <= nums[r] {
            // The smallest number should be the number at left-hand side of m.
            r = m
        }
    }

    return nums[m]   
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
