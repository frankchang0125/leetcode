/*
 * @lc app=leetcode id=241 lang=golang
 *
 * [241] Different Ways to Add Parentheses
 *
 * https://leetcode.com/problems/different-ways-to-add-parentheses/description/
 *
 * algorithms
 * Medium (48.61%)
 * Total Accepted:    68.2K
 * Total Submissions: 139.8K
 * Testcase Example:  '"2-1-1"'
 *
 * Given a string of numbers and operators, return all possible results from
 * computing all the different possible ways to group numbers and operators.
 * The valid operators are +, - and *.
 *
 * Example 1:
 *
 *
 * Input: "2-1-1"
 * Output: [0, 2]
 * Explanation:
 * ((2-1)-1) = 0
 * (2-(1-1)) = 2
 *
 * Example 2:
 *
 *
 * Input: "2*3-4*5"
 * Output: [-34, -14, -10, -10, 10]
 * Explanation:
 * (2*(3-(4*5))) = -34
 * ((2*3)-(4*5)) = -14
 * ((2*(3-4))*5) = -10
 * (2*((3-4)*5)) = -10
 * (((2*3)-4)*5) = 10
 *
 */
// Reference: https://goo.gl/zD2KeQ
func diffWaysToCompute(input string) []int {
	m := make(map[string][]int)
	return ways(m, input)
}

// Split the input by operator from left to right
// and recursively calculate the sub-answers.
func ways(m map[string][]int, input string) []int {
	// If input is already computed, just return the answers list.
	if result, ok := m[input]; ok {
		return result
	}

	ans := []int{}

	for i := 0; i < len(input); i++ {
		op := input[i]
		if op == '+' || op == '-' || op == '*' {
			// Recursively get the answers list of left part of input.
			left := ways(m, input[:i])
			// Recursively get the answers list of right part of input.
			right := ways(m, input[i+1:])

			var result int

			// Take left answers list and right answers list one-by-one
			// to compute the results (Cartesian Product).
			for _, l := range left {
				for _, r := range right {
					switch op {
					case '+':
						result = l + r
					case '-':
						result = l - r
					case '*':
						result = l * r
					}

					ans = append(ans, result)
				}
			}

			m[input] = ans
		}
	}

	// Input is the plain number, ex: '3', '11', '152'... etc.
	// Add it to answers list.
	if len(ans) == 0 {
		result, _ := strconv.Atoi(input)
		ans = append(ans, result)
		m[input] = ans
	}

	return ans
}
