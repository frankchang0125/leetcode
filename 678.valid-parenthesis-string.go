/*
 * @lc app=leetcode id=678 lang=golang
 *
 * [678] Valid Parenthesis String
 *
 * https://leetcode.com/problems/valid-parenthesis-string/description/
 *
 * algorithms
 * Medium (31.60%)
 * Total Accepted:    23.8K
 * Total Submissions: 74.4K
 * Testcase Example:  '"()"'
 *
 *
 * Given a string containing only three types of characters: '(', ')' and '*',
 * write a function to check whether this string is valid. We define the
 * validity of a string by these rules:
 *
 * Any left parenthesis '(' must have a corresponding right parenthesis ')'.
 * Any right parenthesis ')' must have a corresponding left parenthesis '('.
 * Left parenthesis '(' must go before the corresponding right parenthesis ')'.
 * '*' could be treated as a single right parenthesis ')' or a single left
 * parenthesis '(' or an empty string.
 * An empty string is also valid.
 *
 *
 *
 * Example 1:
 *
 * Input: "()"
 * Output: True
 *
 *
 *
 * Example 2:
 *
 * Input: "(*)"
 * Output: True
 *
 *
 *
 * Example 3:
 *
 * Input: "(*))"
 * Output: True
 *
 *
 *
 * Note:
 *
 * The string size will be in the range [1, 100].
 *
 *
 */
// Reference: http://bit.ly/2VmjWyw
func checkValidString(s string) bool {
	// Minimum possible open '('
	// i.e. Change all '*' to ')'
	// 	ex: "(((***)" = "(((())))", minOp = -1.
	minOp := 0

	// Maximum possible open '('
	// i.e. Change all '*' to '('
	//	ex: "(((***)" = "((((()", maxOp = 5.
	maxOp := 0

	for _, c := range s {
		if c == '(' {
			minOp++
		} else {
			// '*' or ')'
			minOp--
		}

		if c == ')' {
			maxOp--
		} else {
			// '(' or '*'
			maxOp++
		}

		// Maximum possible open '(' should always bigger than 0.
		// Otherwise, after we change all '*' to '(', there will still
		// be extra unpaired ')' remained, which is invalid.
		//	ex: "((*))))", maxOp = -1, string changed to "((())))", still invalid.
		if maxOp < 0 {
			return false
		}

		// If minimum possible open '(' is smaller than zero,
		// reset it to zero.
		//
		// If minOp == 0, all '*' are changed to ')' to make string
		// valid parenthesis.
		// 	ex: "(((**)", minOp = 0, string can be changed to: "((()))", valid.
		//
		// If minOp < 0, all extra '*' are changed to 'empty string' to
		// make string valid.
		//	ex: "(((***)", minOp = -1, string can be changed to: "(((_)))", valid.
		//
		// However, if minOp > 0, which means there must be extra '(' remained
		// even we've changed all '*' to ')'.
		//	ex: "((((**)", minOp = 1, string changed to "(((()))", still invalid.
		minOp = max(0, minOp)
	}

	// At last, minimum possible open '(' should be zero.
	// If minOp > 0, which means there must be extra '(' remained
	// even we've changed all '*' to ')'.
	//	ex: "((((**)", minOp = 1, string changed to "(((()))", still invalid.
	if minOp == 0 {
		return true
	}
	return false
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
