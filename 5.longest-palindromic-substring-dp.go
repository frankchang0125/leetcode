/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (26.27%)
 * Total Accepted:    458.7K
 * Total Submissions: 1.7M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: "cbbd"
 * Output: "bb"
 *
 *
 */
// dp[i][j]:
//	If true: s[i->j] is a palindrome string.
//	If false: s[i->j] is not a palindrome string.
//
// s[i->j] is a palindrome string if:
//	1. s[i] == s[j]
//	2. Inner substring of s[i->j]: s[i+1->j-1] is also a
//	   palindrome string.
//
// 	   i.e.:
//		* Inner substring s[i+1->j-1] is also a
//	      palindrome string, or
//		* Length of s[i+1->j-1] is a single character.
//	   	  (i.e. j - i <= 2)
//
// Time complexity: (O^2)
// Space complexity: (O^2)
//
// Reference: http://bit.ly/2W9VOzB
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, len(s))
	for i := range s {
		dp[i] = make([]bool, len(s))
	}

	maxLen := 1
	maxPalindromeStr := string(s[0])

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			isInnerPalindrome := dp[i+1][j-1] || j-i <= 2

			if s[i] == s[j] && isInnerPalindrome {
				dp[i][j] = true

				// Update max palindrome string if substring is a
				// palindrome string with longer length.
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					maxPalindromeStr = s[i : j+1]
				}
			}
		}
	}

	return maxPalindromeStr
}
