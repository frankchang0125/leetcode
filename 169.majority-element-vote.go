/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (52.50%)
 * Likes:    1614
 * Dislikes: 141
 * Total Accepted:    379.2K
 * Total Submissions: 722K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 * 
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,2,3]
 * Output: 3
 * 
 * Example 2:
 * 
 * 
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 * 
 * 
 */
// Vote algorithm:
//  Count the votes ('count') which current number got.
//  If encountering different number, decrease the votes ('count') by 1.
//  If encountering the same number, increase the votes ('count') by 1.
//  If votes ('count') is zero, change winner number to the new number.
//  The answer will be the winner number who survives to the last.
func majorityElement(nums []int) int {
    var cur int
    count := 0

    for _, num := range nums {
        if count == 0 {
            cur = num
            count++
        } else if num == cur {
            count++
        } else if num != cur {
            count--
        }
    }

    return cur
}
