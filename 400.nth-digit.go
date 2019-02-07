/*
 * @lc app=leetcode id=400 lang=golang
 *
 * [400] Nth Digit
 *
 * https://leetcode.com/problems/nth-digit/description/
 *
 * algorithms
 * Easy (30.05%)
 * Total Accepted:    43.8K
 * Total Submissions: 145.7K
 * Testcase Example:  '3'
 *
 * Find the nth digit of the infinite integer sequence 1, 2, 3, 4, 5, 6, 7, 8,
 * 9, 10, 11, ...
 *
 * Note:
 * n is positive and will fit within the range of a 32-bit signed integer (n <
 * 231).
 *
 *
 * Example 1:
 *
 * Input:
 * 3
 *
 * Output:
 * 3
 *
 *
 *
 * Example 2:
 *
 * Input:
 * 11
 *
 * Output:
 * 0
 *
 * Explanation:
 * The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a
 * 0, which is part of the number 10.
 *
 *
 */
func findNthDigit(n int) int {
    i := 0
    base := 1

    for {
        if n-(9*int(math.Pow(10, float64(i)))*(i+1)) > 0 {
            n = n - (9 * int(math.Pow(10, float64(i))) * (i + 1))
            i++
            base = base * 10
        } else {
            break
        }
    }

    var num int

    if n%(i+1) == 0 {
        num = base + (n / (i + 1)) - 1
    } else {
        num = base + (n / (i + 1))
    }

    offset := n%(i+1) - 1
    if offset < 0 {
        offset = i
    }

    numStr := strconv.FormatInt(int64(num), 10)
    result, _ := strconv.ParseInt(string(numStr[offset]), 10, 32)

    return int(result)
}
