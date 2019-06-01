/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 *
 * https://leetcode.com/problems/decode-ways/description/
 *
 * algorithms
 * Medium (21.70%)
 * Total Accepted:    247K
 * Total Submissions: 1.1M
 * Testcase Example:  '"12"'
 *
 * A message containing letters from A-Z is being encoded to numbers using the
 * following mapping:
 *
 *
 * 'A' -> 1
 * 'B' -> 2
 * ...
 * 'Z' -> 26
 *
 *
 * Given a non-empty string containing only digits, determine the total number
 * of ways to decode it.
 *
 * Example 1:
 *
 *
 * Input: "12"
 * Output: 2
 * Explanation: It could be decoded as "AB" (1 2) or "L" (12).
 *
 *
 * Example 2:
 *
 *
 * Input: "226"
 * Output: 3
 * Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2
 * 6).
 *
 */
// References:
//  http://bit.ly/2Vne3RH
//  http://bit.ly/2FDh5v8
func numDecodings(s string) int {
    N := len(s)

    if N == 0 || s[0] == '0' {
		return 0
    }
    
    if N == 1 {
        return 1
    }

    // dp[i]: Ways to decode s[0->i].
    // dp[i] = 0, if: s[i] & s[i-1]s[i] are both NOT valid.
    //  ex: "30"
    // dp[i] = dp[i-1] + dp[i-2], if: s[i] & s[i-1]s[i] are both valid.
    //  ex: "24"
    // dp[i] = dp[i-1], if: s[i] is valid but s[i-1]s[i] is NOT valid.
    //  ex: "27"
    // dp[i] = dp[i-2], if: s[i] is NOT valid but s[i-1]s[i] is valid.
    //  ex: "20"
    //
    // For each iteration, we only use dp[i-1] & dp[i-2].
    // Thus, we can use rolling array technique to keep the decode ways
    // of previous two iterations:
	// 	dp[0]: Ways to decode s[0->i-1].
	//	dp[1]: Ways to decode s[0->i-2].
    dp := make([]int, 2)
    dp[0] = 1

    // Handle dp[1].
    if N >= 2 {
        var w int

        if !isValidChar(s[1]) && !isValidString(s[0], s[1]) {
            w = 0
        } else if isValidChar(s[1]) && isValidString(s[0], s[1]) {
            w = 2
        } else {
            w = 1
        }

        dp[1] = dp[0]
        dp[0] = w
    }

	for i := 2; i < len(s); i++ {
        if !isValidChar(s[i]) && !isValidString(s[i-1], s[i]) {
            // Both s[i] & s[i-1]s[i] are NOT valid.
            return 0
        }
        
        var w int

        if isValidChar(s[i]) {
            // s[i] is valid.
            w = dp[0]
        }

        if isValidString(s[i-1], s[i]) {
            // s[i-1]s[i] is valid.
            w += dp[1]
        }

        dp[1] = dp[0]
        dp[0] = w
	}

	return dp[0]
}

func isValidChar(c byte) bool {
    if c != '0' {
        return true
    }
    return false
}

func isValidString(c1 byte, c2 byte) bool {
    num := (c1 - '0') * 10 + (c2 - '0')
    if num >= 10 && num <= 26 {
        return true
    }
    return false
}
