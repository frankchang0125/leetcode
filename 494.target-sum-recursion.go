/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 *
 * https://leetcode.com/problems/target-sum/description/
 *
 * algorithms
 * Medium (44.74%)
 * Total Accepted:    89.1K
 * Total Submissions: 197.9K
 * Testcase Example:  '[1,1,1,1,1]\n3'
 *
 * 
 * You are given a list of non-negative integers, a1, a2, ..., an, and a
 * target, S. Now you have 2 symbols + and -. For each integer, you should
 * choose one from + and - as its new symbol.
 * ⁠
 * 
 * Find out how many ways to assign symbols to make sum of integers equal to
 * target S.  
 * 
 * 
 * Example 1:
 * 
 * Input: nums is [1, 1, 1, 1, 1], S is 3. 
 * Output: 5
 * Explanation: 
 * 
 * -1+1+1+1+1 = 3
 * +1-1+1+1+1 = 3
 * +1+1-1+1+1 = 3
 * +1+1+1-1+1 = 3
 * +1+1+1+1-1 = 3
 * 
 * There are 5 ways to assign symbols to make the sum of nums be target 3.
 * 
 * 
 * 
 * Note:
 * 
 * The length of the given array is positive and will not exceed 20. 
 * The sum of elements in the given array will not exceed 1000.
 * Your output answer is guaranteed to be fitted in a 32-bit integer.
 * 
 * 
 */
// References:
//  http://bit.ly/2D17b5g
//  http://bit.ly/2CYzBwJ
func findTargetSumWays(nums []int, S int) int {
    return ways(nums, S, 0, 0)
}

func ways(nums []int, target int, sum int, i int) int {
    if i == len(nums) {
        if sum == target {
            return 1
        }
        return 0
    }

    w := ways(nums, target, sum + nums[i], i+1)
    w += ways(nums, target, sum - nums[i], i+1)

    return w
}
