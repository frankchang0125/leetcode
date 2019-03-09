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
	bfs(&ans, digits)
	return ans
}

func bfs(ans *[]string, digits string) {
	queue := []string{}

	digit := digits[0]
	button := buttons[digit]

	for _, c := range button {
		queue = append(queue, string(c)) // Enqueue
	}
	level := 1

	for {
		if level == len(digits) {
			for _, s := range queue {
				*ans = append(*ans, s)
			}
			break
		}

		level++

		for i := len(queue); i > 0; i-- {
			s := queue[0] // Dequeue
			queue = queue[1:]

			for _, c := range buttons[digits[level-1]] {
				queue = append(queue, s + string(c)) // Enqueue
			}
		}
	}
}
