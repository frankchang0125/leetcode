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
	// dp0[i]: Longest longest palindromic subsequence's length
	//         of substring: s[i->j] with length: 'l',
	//		   where j = i + l - 1.
	// dp1[i]: Longest longest palindromic subsequence's length
	//         of substring: s[i->j] with length: 'l - 1',
	//		   where j = i + (l - 1) - 1).
	// dp2[i]: Longest longest palindromic subsequence's length
	//         of substring: s[i->j] with length: 'l - 2',
	//		   where j = i + (l - 2) - 1).
	//
	// If i == j: length = 1, i.e. dp0[i] = 1
	//
	// If s[i] == s[j]:
	//	dp0 = dp2 + 2 (with s[i] and s[j])
	// If s[i] != s[j]:
	//	dp0 = max(dp1[i+1], dp1[i])
	//
	// Ex: "a****a", dp0 = dp2 + 2
	// Ex: "a***ab" or "ab***b", dp0 = max(dp1[i+1], dp1[i])
	N := len(s)
	dp0 := make([]int, N)
	dp1 := make([]int, N)
	dp2 := make([]int, N)

	for l := 1; l <= N; l++ {
		for i := 0; i <= N-l; i++ {
			j := i + l - 1

			if i == j {
				dp0[i] = 1
				continue
			}

			// We solve DP from smaller-length substrings
			// to larger-length substrings.
			// Thus, for dp0[i]: dp2[i+1], dp1[i+1] and dp1[i]
			// are already been solved before.
			if s[i] == s[j] {
				dp0[i] = dp2[i+1] + 2
			} else {
				dp0[i] = max(dp1[i+1], dp1[i])
			}
		}

		dp0, dp1 = dp1, dp0
		dp0, dp2 = dp2, dp0
	}

	return dp1[0]
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
