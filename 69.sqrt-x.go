/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (30.42%)
 * Total Accepted:    350.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 *
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 *
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since
 * the decimal part is truncated, 2 is returned.
 *
 *
 */
// Reference: http://bit.ly/2GhVloe
func mySqrt(x int) int {
	l := 0
	r := x + 1

	for l < r {
		m := l + (r-l)/2

		// Upper bound binary search:
		//  Try to find smallest number: m, which m^2 > x.
		//  i.e. m > sqrt(x)
		if m*m > x {
			// Search left part.
			r = m
		} else {
			// Search right part.
			l = m + 1
		}
	}

	// 'l' is the smallest number which l > sqrt(x),
	// thus, answer is 'l - 1'.
	return l - 1
}
