/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.28%)
 * Likes:    2155
 * Dislikes: 3272
 * Total Accepted:    691.5K
 * Total Submissions: 2.7M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 * 
 * Example 1:
 * 
 * 
 * Input: 123
 * Output: 321
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -123
 * Output: -321
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 120
 * Output: 21
 * 
 * 
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 0 when the reversed
 * integer overflows.
 * 
 */
func reverse(x int) int {
    num := abs(x)
    ans := 0

    for num > 0 {
        ans = ans * 10 + num % 10
        if ans > math.MaxInt32 {
            return 0
        }
        num /= 10
    }

    if x < 0 {
        return -int(ans)
    }

    return int(ans)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
