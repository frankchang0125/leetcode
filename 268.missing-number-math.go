/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 *
 * https://leetcode.com/problems/missing-number/description/
 *
 * algorithms
 * Easy (48.28%)
 * Likes:    880
 * Dislikes: 1317
 * Total Accepted:    274.9K
 * Total Submissions: 568.5K
 * Testcase Example:  '[3,0,1]'
 *
 * Given an array containing n distinct numbers taken from 0, 1, 2, ..., n,
 * find the one that is missing from the array.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,0,1]
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [9,6,4,2,3,5,7,0,1]
 * Output: 8
 * 
 * 
 * Note:
 * Your algorithm should run in linear runtime complexity. Could you implement
 * it using only constant extra space complexity?
 */
// For n distinct numbers, range: [0->n], there
func missingNumber(nums []int) int {
    N := len(nums)

    sum := (1+N)*N/2

    for _, num := range nums {
        sum -= num
    }

    return sum
}
