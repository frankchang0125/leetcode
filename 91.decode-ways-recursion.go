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
	mem := map[string]int{}
	return ways(s, mem)
}

func ways(s string, mem map[string]int) int {
	if len(s) == 0 {
		return 1
	}

	if m, ok := mem[s]; ok {
		return m
	}

	// String cannot starts with '0'.
	if s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	// 1 character decoded.
	w := ways(s[1:], mem)
	mem[s[1:]] = w

	// Check if we can decode first-two characters.
	n, _ := strconv.Atoi(s[:2])
	if n >= 1 && n <= 26 {
		// 2 characters decoded.
		w += ways(s[2:], mem)
		mem[s[2:]] = w
	}

	return w
}
