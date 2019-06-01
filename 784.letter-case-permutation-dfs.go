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
    // Use DFS to generate all letter permutations of string: S.
    // Recursively call to transform one character at a time.
	dfs([]byte(S), 0, &ans)
	return ans
}

// s is the transformed string.
// i is the index of character to be transformed.
func dfs(s []byte, i int, ans *[]string) {
	if i == len(s) {
        // All characters have been transformed, add it to answers.
		*ans = append(*ans, string(s))
		return
	}

    // Original string.
	dfs(s, i+1, ans)

    // Transform string, if character at index: i is an alphabet.
	if isAlphabet(s[i]) {
		s[i] ^= (1 << 5)
		dfs(s, i+1, ans)
	}
}

func isAlphabet(c byte) bool {
    if (c >= 65 && c <= 90) || (c >= 97 && c <= 122 ) {
		return true
	}
    return false
}
