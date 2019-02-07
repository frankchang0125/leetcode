/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 *
 * https://leetcode.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (43.08%)
 * Total Accepted:    344K
 * Total Submissions: 798.2K
 * Testcase Example:  '2'
 *
 * You are climbing a stair case. It takes n steps to reach to the top.
 *
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can
 * you climb to the top?
 *
 * Note: Given n will be a positive integer.
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: 2
 * Explanation: There are two ways to climb to the top.
 * 1. 1 step + 1 step
 * 2. 2 steps
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: 3
 * Explanation: There are three ways to climb to the top.
 * 1. 1 step + 1 step + 1 step
 * 2. 1 step + 2 steps
 * 3. 2 steps + 1 step
 *
 *
 */
 // Reference: https://goo.gl/hZzkqc
func climbStairs(n int) int {
    // Solved with dynamic programming
    dp1, dp2 := 1, 1
    for i := 2; i <= n; i++ {
        dp2, dp1 = dp1+dp2, dp2
    }
    return dp2
}
