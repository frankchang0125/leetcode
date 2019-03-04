/*
 * @lc app=leetcode id=680 lang=golang
 *
 * [680] Valid Palindrome II
 *
 * https://leetcode.com/problems/valid-palindrome-ii/description/
 *
 * algorithms
 * Easy (33.56%)
 * Total Accepted:    62K
 * Total Submissions: 183.9K
 * Testcase Example:  '"aba"'
 *
 *
 * Given a non-empty string s, you may delete at most one character.  Judge
 * whether you can make it a palindrome.
 *
 *
 * Example 1:
 *
 * Input: "aba"
 * Output: True
 *
 *
 *
 * Example 2:
 *
 * Input: "abca"
 * Output: True
 * Explanation: You could delete the character 'c'.
 *
 *
 *
 * Note:
 *
 * The string will only contain lowercase characters a-z.
 * The maximum length of the string is 50000.
 *
 *
 */
// Reference: https://goo.gl/xPCZRs
func validPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	// Compare the characters from both sides to middle.
	l := 0
	r := len(s) - 1

	for {
		if l >= r {
			break
		}

		// By deleting at most one character to make string a palindrome,
		// when a non-palindrome character is found, the string can still be
		// a palindrome if:
		//	1. Delete character at 'l' index, i.e. remaining string: s[l+1, r]
		//	   is still a palindrome.
		//  2. Delete character at 'r' index, i.e. remaining string: s[l, r-1]
		//	   is still a palindrome.
		if s[r] != s[l] {
			return isPalindrome(s, l+1, r) ||
				isPalindrome(s, l, r-1)
		}

		l++
		r--
	}

	return true
}

func isPalindrome(s string, l, r int) bool {
	for {
		if l >= r {
			break
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
