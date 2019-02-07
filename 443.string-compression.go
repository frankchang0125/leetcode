/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 *
 * https://leetcode.com/problems/string-compression/description/
 *
 * algorithms
 * Easy (36.57%)
 * Total Accepted:    40.5K
 * Total Submissions: 110.9K
 * Testcase Example:  '["a","a","b","b","c","c","c"]'
 *
 * Given an array of characters, compress it in-place.
 *
 * The length after compression must always be smaller than or equal to the
 * original array.
 *
 * Every element of the array should be a character (not int) of length 1.
 *
 * After you are done modifying the input array in-place, return the new length
 * of the array.
 *
 *
 * Follow up:
 * Could you solve it using only O(1) extra space?
 *
 *
 * Example 1:
 *
 *
 * Input:
 * ["a","a","b","b","c","c","c"]
 *
 * Output:
 * Return 6, and the first 6 characters of the input array should be:
 * ["a","2","b","2","c","3"]
 *
 * Explanation:
 * "aa" is replaced by "a2". "bb" is replaced by "b2". "ccc" is replaced by
 * "c3".
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ["a"]
 *
 * Output:
 * Return 1, and the first 1 characters of the input array should be: ["a"]
 *
 * Explanation:
 * Nothing is replaced.
 *
 *
 *
 *
 * Example 3:
 *
 *
 * Input:
 * ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
 *
 * Output:
 * Return 4, and the first 4 characters of the input array should be:
 * ["a","b","1","2"].
 *
 * Explanation:
 * Since the character "a" does not repeat, it is not compressed.
 * "bbbbbbbbbbbb" is replaced by "b12".
 * Notice each digit has it's own entry in the array.
 *
 *
 *
 *
 * Note:
 *
 *
 * All characters have an ASCII value in [35, 126].
 * 1 <= len(chars) <= 1000.
 *
 *
 */
func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	} else if len(chars) == 1 {
		return 1
	}

	comp := chars[0]
	count := 1
	compressIdx := 0

	for i := 1; i <= len(chars); i++ {
		if i < len(chars) && chars[i] == comp {
			count++
		} else {
			// chars[i] != comp or i == len(chars)
			chars[compressIdx] = comp
			compressIdx++

			if count > 1 {
				countStr := strconv.Itoa(count)
				for j := 0; j < len(countStr); j++ {
					chars[compressIdx+j] = countStr[j]
				}
				compressIdx += len(countStr)
			}

			if i < len(chars) {
				comp = chars[i]
			}
			count = 1
		}
	}

	return compressIdx
}
