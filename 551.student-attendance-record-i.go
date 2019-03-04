/*
 * @lc app=leetcode id=551 lang=golang
 *
 * [551] Student Attendance Record I
 *
 * https://leetcode.com/problems/student-attendance-record-i/description/
 *
 * algorithms
 * Easy (45.14%)
 * Total Accepted:    47.2K
 * Total Submissions: 104.6K
 * Testcase Example:  '"PPALLP"'
 *
 * You are given a string representing an attendance record for a student. The
 * record only contains the following three characters:
 *
 *
 *
 * 'A' : Absent.
 * 'L' : Late.
 * ⁠'P' : Present.
 *
 *
 *
 *
 * A student could be rewarded if his attendance record doesn't contain more
 * than one 'A' (absent) or more than two continuous 'L' (late).
 *
 * You need to return whether the student could be rewarded according to his
 * attendance record.
 *
 * Example 1:
 *
 * Input: "PPALLP"
 * Output: True
 *
 *
 *
 * Example 2:
 *
 * Input: "PPALLL"
 * Output: False
 *
 *
 *
 *
 *
 */
// Reference: https://goo.gl/xr3sGf
func checkRecord(s string) bool {
	if len(s) < 2 {
		return true
	}

	absentCount := 0
	lateContCount := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			absentCount++
		}

		if s[i] == 'L' {
			lateContCount++
		} else {
			lateContCount = 0
		}

		if absentCount > 1 || lateContCount > 2 {
			return false
		}
	}

	return true
}
