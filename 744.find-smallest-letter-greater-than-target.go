/*
 * @lc app=leetcode id=744 lang=golang
 *
 * [744] Find Smallest Letter Greater Than Target
 *
 * https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/
 *
 * algorithms
 * Easy (43.45%)
 * Total Accepted:    37.1K
 * Total Submissions: 84.8K
 * Testcase Example:  '["c","f","j"]\n"a"'
 *
 *
 * Given a list of sorted characters letters containing only lowercase letters,
 * and given a target letter target, find the smallest element in the list that
 * is larger than the given target.
 *
 * Letters also wrap around.  For example, if the target is target = 'z' and
 * letters = ['a', 'b'], the answer is 'a'.
 *
 *
 * Examples:
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "a"
 * Output: "c"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "c"
 * Output: "f"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "d"
 * Output: "f"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "g"
 * Output: "j"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "j"
 * Output: "c"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "k"
 * Output: "c"
 *
 *
 *
 * Note:
 *
 * letters has a length in range [2, 10000].
 * letters consists of lowercase letters, and contains at least 2 unique
 * letters.
 * target is a lowercase letter.
 *
 *
 */
// Reference: http://bit.ly/2KoScYW
func nextGreatestLetter(letters []byte, target byte) byte {
    // Left included + Right not included.
    //
	// Because we want to find the smallest number
	// which is greater than the target, thus Right is not included.
	// Therefore, at the end of the loop, 'l' will equal to 'r',
    // which will be the the index of smallest number greater than target
    // (wrap around if needed).
	l := 0
	r := len(letters)

	for {
		if l >= r {
			break
		}

		i := l + (r-l)/2

		if target >= letters[i] {
			// Search right part.
			l = i + 1
		} else {
			// Search left part.
			r = i
		}
	}

	return letters[l%len(letters)]
}
