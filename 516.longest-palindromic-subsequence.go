/*
 * @lc app=leetcode id=516 lang=golang
 *
 * [516] Longest Palindromic Subsequence
 *
 * https://leetcode.com/problems/longest-palindromic-subsequence/description/
 *
 * algorithms
 * Medium (44.99%)
 * Total Accepted:    53.6K
 * Total Submissions: 117.5K
 * Testcase Example:  '"bbbab"'
 *
 *
 * Given a string s, find the longest palindromic subsequence's length in s.
 * You may assume that the maximum length of s is 1000.
 *
 *
 * Example 1:
 * Input:
 *
 * "bbbab"
 *
 * Output:
 *
 * 4
 *
 * One possible longest palindromic subsequence is "bbbb".
 *
 *
 * Example 2:
 * Input:
 *
 * "cbbd"
 *
 * Output:
 *
 * 2
 *
 * One possible longest palindromic subsequence is "bb".
 *
 */
// Reference: https://goo.gl/KBVbkB
func longestPalindromeSubseq(s string) int {
	// dp[i][j]: Longest longest palindromic subsequence's length
	//			 of substring: s[i->j].
	// dp[i][i] = 1
	//
	// If s[i] == s[j]:
	//	dp[i][j] = dp[i-1][j-1] + 2 (with s[i] and s[j])
	// If s[i] != s[j]:
	//	dp[i][j] = max(dp[i][j-1], dp[i+1][j])
	//
	// Ex: "a****a", dp[i][j] = dp[i+1][j-1] + 2.
	// Ex: "a***ab", dp[i][j] = dp[i][j-1]
	//	   (Note: s[j-1] doesn't need to be 'a' exactly.)
	// Ex: "ab***b", dp[i][j] = dp[i+1][j]
	//	   (Note: s[i+1] doesn't need to be 'b' exactly.)
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}

	for l := 1; l <= len(s); l++ {
		for i := 0; i <= len(s)-l; i++ {
			j := i + l - 1

			if i == j {
				dp[i][j] = 1
				continue
			}

			// We solve DP from smaller-length substrings
			// to larger-length substrings.
			// Thus, for dp[i][j]: dp[i+1][j-1], dp[i][j-1] and dp[i+1][j]
			// are already been solved before.
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][len(s)-1]
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
