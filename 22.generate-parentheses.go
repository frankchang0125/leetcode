/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (52.61%)
 * Total Accepted:    291.7K
 * Total Submissions: 554.2K
 * Testcase Example:  '3'
 *
 *
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 *
 *
 *
 * For example, given n = 3, a solution set is:
 *
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
 */
func generateParenthesis(n int) []string {
	ans := []string{}
	parentheses(n, n, []byte{}, &ans)
	return ans
}

func parentheses(left, right int, p []byte, ans *[]string) {
    // Valid well-formed parentheses:
    //  Remaining left count cannot bigger than remaining right count.
	if left > right || left < 0 || right < 0 {
		return
	}

	if left == 0 && right == 0 {
		*ans = append(*ans, string(p))
		return
	}

	parentheses(left-1, right, append(p, '('), ans)
	parentheses(left, right-1, append(p, ')'), ans)
}
