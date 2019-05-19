/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (27.92%)
 * Likes:    826
 * Dislikes: 2045
 * Total Accepted:    316.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (x^n).
 * 
 * Example 1:
 * 
 * 
 * Input: 2.00000, 10
 * Output: 1024.00000
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 2.10000, 3
 * Output: 9.26100
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 2.00000, -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 * 
 * 
 * Note:
 * 
 * 
 * -100.0 < x < 100.0
 * n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
 * 
 * 
 */
func myPow(x float64, n int) float64 {
    result := power(x, n)
    if n < 0 {
        return 1 / result
    }
    return result
}

func power(x float64, n int) float64 {
    if n == 0 {
        return 1
    }

    if x == 1 {
        return 1
    }
    
    if x == -1 {
        if n & 1 == 0 {
            return 1
        } else {
            return -1
        }
    }

    exp := abs(n)
    result := x
    for i := 0; i < exp-1; i++ {
        result *= x
    }
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
