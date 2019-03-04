/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 *
 * https://leetcode.com/problems/longest-palindrome/description/
 *
 * algorithms
 * Easy (47.23%)
 * Total Accepted:    87.3K
 * Total Submissions: 184.4K
 * Testcase Example:  '"abccccdd"'
 *
 * Given a string which consists of lowercase or uppercase letters, find the
 * length of the longest palindromes that can be built with those letters.
 *
 * This is case sensitive, for example "Aa" is not considered a palindrome
 * here.
 *
 * Note:
 * Assume the length of given string will not exceed 1,010.
 *
 *
 * Example:
 *
 * Input:
 * "abccccdd"
 *
 * Output:
 * 7
 *
 * Explanation:
 * One longest palindrome that can be built is "dccaccd", whose length is 7.
 *
 *
 */
// Reference: https://goo.gl/19pjry
func longestPalindrome(s string) int {
	// String length less than 2 must be palindrome.
	if len(s) < 2 {
		return len(s)
	}

	// Preprocessing, count the number of characters of string.
	charCounts := make([]int, 128)
	for _, c := range s {
		charCounts[c]++
	}

	// We allow odd number of characters to appear once only.
	// If more than one character is odd nubmer, set 'odd' flag to true,
	// which means we can take one more character as the middle character
	// of the longest palindrome.
	odd := false
	palindromeLength := 0

	// Count the length of longest palindrome.
	for _, count := range charCounts {
		palindromeLength += 2 * (count / 2)
		if !odd && count%2 != 0 {
			odd = true
		}
	}

	if odd {
		// Take one more character as the middle character
		// of the longest palindrome.
		palindromeLength++
	}

	return palindromeLength
}
