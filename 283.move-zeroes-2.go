/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 *
 * https://leetcode.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (54.23%)
 * Likes:    2019
 * Dislikes: 73
 * Total Accepted:    466.8K
 * Total Submissions: 859.2K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * Given an array nums, write a function to move all 0's to the end of it while
 * maintaining the relative order of the non-zero elements.
 * 
 * Example:
 * 
 * 
 * Input: [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 * 
 * Note:
 * 
 * 
 * You must do this in-place without making a copy of the array.
 * Minimize the total number of operations.
 * 
 */
// Copy all nonzero numbers from the front of array,
// at the same time, find out the first zero's index
// which zeroes should be placed.
// Then, since we know first zero's index, we can fill the rest of
// array with zeroes.
//
// ex: [0,1,0,3,12]
//  Step 1. Copy all nonzero numbers from the front of array:
//          => [1,3,12,3,12], first zero's index: 'i' will be: 3
//  Step 2. Fill the rest of array with zeroes:
//          => [1,3,12,0,0]
func moveZeroes(nums []int)  {
    // i will eventually be the index of first zero's index
    // which zeroes should be placed.
    i := 0

    // Copy all nonzero numbers from the front of array.
    for _, v := range nums {
        if v != 0 {
            nums[i] = v
            i++
        }
    }

    // i is not the first zero's index, fill the rest of array
    // with zeros.
    for j := i; j < len(nums); j++ {
        nums[j] = 0
    }
}
