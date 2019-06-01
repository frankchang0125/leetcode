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
func moveZeroes(nums []int)  {
    N := len(nums)
    if N <= 1 {
        return
    }
    
    zero := 0
    nonzero := 0
    
    for {
        // Find next zero's index.
        for zero < N {
            if nums[zero] == 0 {
                break
            }
            zero++
        }
        
        // Find next nonzero number's index.
        for nonzero < N {
            if nums[nonzero] != 0 {
                break
            }
            nonzero++
        }
        
        // Boundary check.
        if zero >= N || nonzero >= N {
            break
        }
        
        // If nonzero number is after zero, swap them.
        if nonzero > zero {
            nums[zero], nums[nonzero] = nums[nonzero], nums[zero]
        }

        // Update nonzero index to compare more numbers.
        nonzero = zero
    } 
}
