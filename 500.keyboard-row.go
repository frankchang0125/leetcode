/*
 * @lc app=leetcode id=500 lang=golang
 *
 * [500] Keyboard Row
 *
 * https://leetcode.com/problems/keyboard-row/description/
 *
 * algorithms
 * Easy (61.47%)
 * Total Accepted:    80.8K
 * Total Submissions: 131.5K
 * Testcase Example:  '["Hello","Alaska","Dad","Peace"]'
 *
 * Given a List of words, return the words that can be typed using letters of
 * alphabet on only one row's of American keyboard like the image
 * below.
 *
 *
 *
 *
 *
 *
 * Example:
 *
 *
 * Input: ["Hello", "Alaska", "Dad", "Peace"]
 * Output: ["Alaska", "Dad"]
 *
 *
 *
 *
 * Note:
 *
 *
 * You may use one character in the keyboard more than once.
 * You may assume the input string will only contain letters of alphabet.
 *
 *
 */
func findWords(words []string) []string {
	row1 := map[rune]bool{
		'q': true,
		'w': true,
		'e': true,
		'r': true,
		't': true,
		'y': true,
		'u': true,
		'i': true,
		'o': true,
		'p': true,
	}

	row2 := map[rune]bool{
		'a': true,
		's': true,
		'd': true,
		'f': true,
		'g': true,
		'h': true,
		'j': true,
		'k': true,
		'l': true,
	}

	row3 := map[rune]bool{
		'z': true,
		'x': true,
		'c': true,
		'v': true,
		'b': true,
		'n': true,
		'm': true,
	}

	rows := []map[rune]bool{
		row1,
		row2,
		row3,
	}

	result := make([]string, 0)

	for _, word := range words {
		for _, row := range rows {
			matchCount := 0
			matches := make(map[rune]bool)

			for _, c := range strings.ToLower(word) {
				if _, ok := matches[c]; ok {
					matchCount++
					continue
				}

				if _, ok := row[c]; ok {
					matchCount++
					matches[c] = true
				} else {
					break
				}
			}

			if matchCount == len(word) {
				result = append(result, word)
				break
			}
		}
	}

	return result
}
