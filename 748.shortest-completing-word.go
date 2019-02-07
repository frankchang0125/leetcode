/*
 * @lc app=leetcode id=748 lang=golang
 *
 * [748] Shortest Completing Word
 *
 * https://leetcode.com/problems/shortest-completing-word/description/
 *
 * algorithms
 * Easy (53.25%)
 * Total Accepted:    16.3K
 * Total Submissions: 30.5K
 * Testcase Example:  '"1s3 PSt"\n["step","steps","stripe","stepple"]'
 *
 *
 * Find the minimum length word from a given dictionary words, which has all
 * the letters from the string licensePlate.  Such a word is said to complete
 * the given string licensePlate
 *
 * Here, for letters we ignore case.  For example, "P" on the licensePlate
 * still matches "p" on the word.
 *
 * It is guaranteed an answer exists.  If there are multiple answers, return
 * the one that occurs first in the array.
 *
 * The license plate might have the same letter occurring multiple times.  For
 * example, given a licensePlate of "PP", the word "pair" does not complete the
 * licensePlate, but the word "supper" does.
 *
 *
 * Example 1:
 *
 * Input: licensePlate = "1s3 PSt", words = ["step", "steps", "stripe",
 * "stepple"]
 * Output: "steps"
 * Explanation: The smallest length word that contains the letters "S", "P",
 * "S", and "T".
 * Note that the answer is not "step", because the letter "s" must occur in the
 * word twice.
 * Also note that we ignored case for the purposes of comparing whether a
 * letter exists in the word.
 *
 *
 *
 * Example 2:
 *
 * Input: licensePlate = "1s3 456", words = ["looks", "pest", "stew", "show"]
 * Output: "pest"
 * Explanation: There are 3 smallest length words that contains the letters
 * "s".
 * We return the one that occurred first.
 *
 *
 *
 * Note:
 *
 * licensePlate will be a string with length in range [1, 7].
 * licensePlate will contain digits, spaces, or letters (uppercase or
 * lowercase).
 * words will have a length in the range [10, 1000].
 * Every words[i] will consist of lowercase letters, and have length in range
 * [1, 15].
 *
 *
 */
func shortestCompletingWord(licensePlate string, words []string) string {
	plate := make(map[rune]int)

	// Preprocessing, count the number of each character appeared in the
	// licence plate
	for _, char := range licensePlate {
		alphabet, upper := isAlphabet(char)
		if alphabet {
			var c = char
			if upper {
				c = c + 32
			}
			if _, ok := plate[c]; !ok {
				plate[c] = 1
			} else {
				plate[c]++
			}
		}
	}

	// Minimum length of word that contains the letters.
	var minLength = 30
	// The index of minimum length of word which has all letters in licence plate.
	var minWordIdx = -1

	for i, word := range words {
		// Word has larger length, no need to compare.
		if len(word) >= minLength {
			continue
		}

		// Count the number of each character appeared in the word.
		charCount := make(map[rune]int)
		for _, c := range word {
			if _, ok := charCount[c]; !ok {
				charCount[c] = 1
			} else {
				charCount[c]++
			}
		}

		// Just iterate every characters in licence plate and make sure
		// the count of each character appeared in word is larger or equals to
		// the count of correspond character in licence plate.
		var found = true
		for char, count := range plate {
			wCount, ok := charCount[char]
			if !ok {
				// Character in licence plate cannot be found in word.
				found = false
				break
			} else if wCount < count {
				// Character count in licence plate is larger than the count
				// in word.
				found = false
				break
			}
		}

		if found {
			minLength = len(word)
			minWordIdx = i
		}
	}

	if minWordIdx == -1 {
        // Nothing matched, return empty string.
		return ""
	}

	return words[minWordIdx]
}

func isAlphabet(c rune) (alphabet bool, upper bool) {
	if c >= 65 && c <= 90 {
		return true, true
	} else if c >= 97 && c <= 122 {
		return true, false
	}
	return false, false
}
