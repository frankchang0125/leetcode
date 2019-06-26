/*
 * @lc app=leetcode id=856 lang=golang
 *
 * [856] Score of Parentheses
 *
 * https://leetcode.com/problems/score-of-parentheses/description/
 *
 * algorithms
 * Medium (55.00%)
 * Total Accepted:    15.6K
 * Total Submissions: 28.3K
 * Testcase Example:  '"()"'
 *
 * Given a balanced parentheses string S, compute the score of the string based
 * on the following rule:
 * 
 * 
 * () has score 1
 * AB has score A + B, where A and B are balanced parentheses strings.
 * (A) has score 2 * A, where A is a balanced parentheses string.
 * 
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "()"
 * Output: 1
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "(())"
 * Output: 2
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "()()"
 * Output: 2
 * 
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: "(()(()))"
 * Output: 6
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * S is a balanced parentheses string, containing only ( and ).
 * 2 <= S.length <= 50
 * 
 * 
 * 
 * 
 * 
 * 
 */
// Reference: https://goo.gl/bjxzfM
func scoreOfParentheses(S string) int {
    return score(S, 0, len(S) - 1)
}

func score(S string, l int, r int) int {
    if r - l == 1 {
        return 1
    }

    b := 0
    for i := l; i < r; i++ {
        if S[i] == '(' {
            b++
        }

        if S[i] == ')' {
            b--
        }

        if b == 0 {
            return score(S, l, i) + score(S, i+1, r)
        }
    }

    return 2 * score(S, l+1, r-1)
}
