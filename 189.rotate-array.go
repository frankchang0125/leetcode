/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 *
 * https://leetcode.com/problems/rotate-array/description/
 *
 * algorithms
 * Easy (28.54%)
 * Likes:    1272
 * Dislikes: 586
 * Total Accepted:    294K
 * Total Submissions: 981.7K
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * Given an array, rotate the array to the right by k steps, where k is
 * non-negative.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,3,4,5,6,7] and k = 3
 * Output: [5,6,7,1,2,3,4]
 * Explanation:
 * rotate 1 steps to the right: [7,1,2,3,4,5,6]
 * rotate 2 steps to the right: [6,7,1,2,3,4,5]
 * rotate 3 steps to the right: [5,6,7,1,2,3,4]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [-1,-100,3,99] and k = 2
 * Output: [3,99,-1,-100]
 * Explanation: 
 * rotate 1 steps to the right: [99,-1,-100,3]
 * rotate 2 steps to the right: [3,99,-1,-100]
 * 
 * 
 * Note:
 * 
 * 
 * Try to come up as many solutions as you can, there are at least 3 different
 * ways to solve this problem.
 * Could you do it in-place with O(1) extra space?
 * 
 */
// Reference: http://bit.ly/2W9RUK9
//  Note: Reference is rotate 'left' k steps.
//        Rotate 'right' is equal to rotate 'left': (len(nums) - k) steps.
func rotate(nums []int, k int)  {
    if len(nums) <= 1 {
        return
    }

    k = k % len(nums)
    j := len(nums) - k

    left := nums[0:j]
    right := nums[j:]
    
    reverse(left)
    reverse(right)
    reverse(nums)
}

func reverse(nums []int) {
    if len(nums) <= 1 {
        return
    }

    l := 0
    r := len(nums) - 1

    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}
