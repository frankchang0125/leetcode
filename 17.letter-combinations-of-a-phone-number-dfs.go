/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (39.96%)
 * Total Accepted:    351.5K
 * Total Submissions: 868K
 * Testcase Example:  ''23''
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent.
 *
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 *
 * Example:
 *
 *
 * Input: '23'
 * Output: ['ad', 'ae', 'af', 'bd', 'be', 'bf', 'cd', 'ce', 'cf'].
 *
 *
 * Note:
 *
 * Although the above answer is in lexicographical order, your answer could be
 * in any order you want.
 *
 */
var buttons = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// Reference: https://goo.gl/nNMzad
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	ans := []string{}
	dfs(&ans, digits, 0, []byte{})
	return ans
}

func dfs(ans *[]string, digits string, level int, cur []byte) {
	if len(cur) == len(digits) {
		c := string(cur)
		*ans = append(*ans, c)
		return
	}

	for i := 0; i < len(buttons[digits[level]]); i++ {
		c := buttons[digits[level]][i]
		cur = append(cur, c)
		dfs(ans, digits, level+1, cur)
		cur = cur[:len(cur)-1]
	}
}
