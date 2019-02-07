/*
 * @lc app=leetcode id=848 lang=golang
 *
 * [848] Shifting Letters
 *
 * https://leetcode.com/problems/shifting-letters/description/
 *
 * algorithms
 * Medium (38.50%)
 * Total Accepted:    9.9K
 * Total Submissions: 25.7K
 * Testcase Example:  '"abc"\n[3,5,9]'
 *
 * We have a string S of lowercase letters, and an integer array shifts.
 *
 * Call the shift of a letter, the next letter in the alphabet, (wrapping
 * around so that 'z' becomes 'a').
 *
 * For example, shift('a') = 'b', shift('t') = 'u', and shift('z') = 'a'.
 *
 * Now for each shifts[i] = x, we want to shift the first i+1 letters of S, x
 * times.
 *
 * Return the final string after all such shifts to S are applied.
 *
 * Example 1:
 *
 *
 * Input: S = "abc", shifts = [3,5,9]
 * Output: "rpl"
 * Explanation:
 * We start with "abc".
 * After shifting the first 1 letters of S by 3, we have "dbc".
 * After shifting the first 2 letters of S by 5, we have "igc".
 * After shifting the first 3 letters of S by 9, we have "rpl", the answer.
 *
 *
 * Note:
 *
 *
 * 1 <= S.length = shifts.length <= 20000
 * 0 <= shifts[i] <= 10 ^ 9
 *
 *
 */
// Reference: https://goo.gl/cJa9PD
func shiftingLetters(S string, shifts []int) string {
    prevShift := 0
    result := make([]byte, len(S))

    for i := len(shifts) - 1; i >= 0; i-- {
        shift := int(S[i]) + shifts[i] + prevShift - 'a'
        shift = shift % 26
        result[i] = byte(shift + 'a')
        prevShift += shifts[i]
    }

    return string(result)
}
