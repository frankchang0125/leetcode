/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (26.03%)
 * Total Accepted:    730.3K
 * Total Submissions: 2.8M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 *
 *
 *
 *
 *
 */
// Use points: 'i' and 'j' to form a sliding window
// of substring: s[i->j-1] without repeating characters.
//
// Use hashmap (or set) to track the non-repeated characters
// in the sliding window: s[i->j-1].
//
// When encountering new non-repeated character:
//  1. Move 'j' right one position to increase the size of sliding window.
//  2. Add the character to the set.
//
// When encountering a repeated character:
//  1. Keep removing the characters starts from position 'i' from hashmap
//     until the repeated character is also been removed and can be
//     re-added to the hashmap.
//     (Note: 'j' is not updated during removes.)
//
// Reference: http://bit.ly/2DoDEmk
func lengthOfLongestSubstring(s string) int {
	// i = 0, j = 0, no character is selected in sliding window,
	// size of sliding window = 0.
	i := 0
	j := 0
	maxLength := 0
	set := make(map[byte]bool)

	for {
		if j == len(s) {
			// All characters have been scanned.
			break
		}

		if _, ok := set[s[j]]; !ok {
			// New non-repeated character can be added to sliding window,
			// add it the the set and move 'j' right one position
			// to increase the size of sliding window.
			set[s[j]] = true
			j++

			length := j - i
			if length > maxLength {
				// Update 'maxLength' if current sliding window size is bigger.
				maxLength = length
			}
		} else {
			// A repeated character in sliding window,
			// keep removing the characters starts from position 'i' from hashmap
			// until the repeated character is also been removed and can be
			// re-added to the hashmap.
			for {
				if _, ok := set[s[j]]; ok {
					delete(set, s[i])
					i++
				} else {
					break
				}
			}
		}
	}

	return maxLength
}
