/*
 * @lc app=leetcode id=784 lang=golang
 *
 * [784] Letter Case Permutation
 *
 * https://leetcode.com/problems/letter-case-permutation/description/
 *
 * algorithms
 * Easy (54.92%)
 * Total Accepted:    36.6K
 * Total Submissions: 66.5K
 * Testcase Example:  '"a1b2"'
 *
 * Given a string S, we can transform every letter individually to be lowercase
 * or uppercase to create another string.  Return a list of all possible
 * strings we could create.
 *
 *
 * Examples:
 * Input: S = "a1b2"
 * Output: ["a1b2", "a1B2", "A1b2", "A1B2"]
 *
 * Input: S = "3z4"
 * Output: ["3z4", "3Z4"]
 *
 * Input: S = "12345"
 * Output: ["12345"]
 *
 *
 * Note:
 *
 *
 * S will be a string with length between 1 and 12.
 * S will consist only of letters or digits.
 *
 *
 */
// Reference: https://goo.gl/YrnF9M
func letterCasePermutation(S string) []string {
	ans := []string{}
	// Use BFS to generate all letter permutations of string: S.
	bfs(S, &ans)
	return ans
}

func bfs(S string, ans *[]string) {
	queue := []string{}
	found := make(map[string]bool)

	queue = append(queue, S) // Enqueue

	// i is the index of character to be transformed.
	i := 0

	for {
		if len(queue) == 0 {
			return
		}

		// Empty the queue of same level.
		// (i.e. All words with same index of character been transformed.)
		for j := len(queue); j > 0; j-- {
			w := queue[0]
			queue = queue[1:] // Dequeue

			// If the word was already found before, skip it.
			if _, ok := found[w]; ok {
				continue
			}

			*ans = append(*ans, w)
			found[w] = true

			for k := i; k < len(w); k++ {
				// Transform string, if character at index: i is an alphabet.
				if isAlphabet(w[k]) {
					t := []byte(w)
					t[k] ^= (1 << 5)
					tStr := string(t)

					if _, ok := found[tStr]; !ok {
						queue = append(queue, tStr) // Enqueue
					}
				}
			}
		}

		// Increase the index of character to be transformed.
		i++
	}
}

func isAlphabet(c byte) bool {
	if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
		return true
	}
	return false
}
