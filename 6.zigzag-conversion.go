/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (30.30%)
 * Total Accepted:    279.5K
 * Total Submissions: 921.7K
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 *
 * And then read line by line: "PAHNAPLSIIGYIR"
 *
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 *
 *
 * string convert(string s, int numRows);
 *
 * Example 1:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 *
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 */
func convert(s string, numRows int) string {
    if numRows == 0 || numRows == 1 {
        return s
    }

    gap := numRows + (numRows - 2)
    result := []byte{}

    for row := 0; row < numRows; row++ {
        i := 0
        j := gap - row*2

        for {
            var idx int

            if row == 0 || row == numRows-1 {
                idx = row + gap*i
                if idx >= len(s) {
                    break
                }
                result = append(result, s[idx])
            } else {
                idx = row + gap*i
                if idx >= len(s) {
                    break
                }
                result = append(result, s[idx])

                idx = row + gap*i + j
                if idx >= len(s) {
                    break
                }
                result = append(result, s[idx])
            }

            i++
        }
    }

    return string(result)
}
