/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 *
 * https://leetcode.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (29.14%)
 * Likes:    2027
 * Dislikes: 89
 * Total Accepted:    207.7K
 * Total Submissions: 712.5K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * Given an integer array nums, find the contiguous subarray within an array
 * (containing at least one number) which has the largest product.
 * 
 * Example 1:
 * 
 * 
 * Input: [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 * 
 */
// Multiply two negative integers will have positive result,
// which may become a larger number than multiplying two positive integers.
// Therefore, we have to maintain two results when using dynamic programming:
//  f[i] = The maximum product of subarray including nums[i].
//  g[i] = The minimum product of subarray including nums[i].
//  f[i] = max(max(f[i-1] * nums[i], g[i-1] * nums[i]), nums[i]))
//  g[i] = min(min(g[i-1] * nums[i], f[i-1] * nums[i]), nums[i]))
//  f[0] = nums[0]
//  g[0] = nums[0]
//
// By using rolling array, we only need to keep:
//  * Maximum product: f
//  * Minimum product: g
// of previous iteration.
//
// Reference: http://bit.ly/2w796kZ
func maxProduct(nums []int) int {
    N := len(nums)
    if N == 0 {
        return 0
    }

    if N == 1 {
        return nums[0]
    }
    
    f := nums[0]
    g := nums[0]
    result := nums[0]

    for i := 1; i < N; i++ {
        maximum := f
        minimum := g
        f = max(max(maximum * nums[i], minimum * nums[i]), nums[i])
        g = min(min(minimum * nums[i], maximum * nums[i]), nums[i])
        result = max(f, result)
    }

    return result
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
