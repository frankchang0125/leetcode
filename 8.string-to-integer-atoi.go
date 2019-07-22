/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 *
 * https://leetcode.com/problems/string-to-integer-atoi/description/
 *
 * algorithms
 * Medium (14.63%)
 * Likes:    936
 * Dislikes: 5921
 * Total Accepted:    365.1K
 * Total Submissions: 2.5M
 * Testcase Example:  '"42"'
 *
 * Implement atoi which converts a string to an integer.
 * 
 * The function first discards as many whitespace characters as necessary until
 * the first non-whitespace character is found. Then, starting from this
 * character, takes an optional initial plus or minus sign followed by as many
 * numerical digits as possible, and interprets them as a numerical value.
 * 
 * The string can contain additional characters after those that form the
 * integral number, which are ignored and have no effect on the behavior of
 * this function.
 * 
 * If the first sequence of non-whitespace characters in str is not a valid
 * integral number, or if no such sequence exists because either str is empty
 * or it contains only whitespace characters, no conversion is performed.
 * 
 * If no valid conversion could be performed, a zero value is returned.
 * 
 * Note:
 * 
 * 
 * Only the space character ' ' is considered as whitespace character.
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. If the numerical
 * value is out of the range of representable values, INT_MAX (2^31 − 1) or
 * INT_MIN (−2^31) is returned.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "42"
 * Output: 42
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "   -42"
 * Output: -42
 * Explanation: The first non-whitespace character is '-', which is the minus
 * sign.
 * Then take as many numerical digits as possible, which gets 42.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "4193 with words"
 * Output: 4193
 * Explanation: Conversion stops at digit '3' as the next character is not a
 * numerical digit.
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: "words and 987"
 * Output: 0
 * Explanation: The first non-whitespace character is 'w', which is not a
 * numerical 
 * digit or a +/- sign. Therefore no valid conversion could be performed.
 * 
 * Example 5:
 * 
 * 
 * Input: "-91283472332"
 * Output: -2147483648
 * Explanation: The number "-91283472332" is out of the range of a 32-bit
 * signed integer.
 * Thefore INT_MIN (−2^31) is returned.
 * 
 */
func myAtoi(str string) int {
    minus := false
    signFound := false

    // Find the start index of the first valid digit.
    var start int
    for start = 0; start < len(str); start++ {
        c := str[start]

        if isWhiteSpace(c) {
            if signFound {
                return 0
            }
            continue
        }
        
        if isPlusSign(c) {
            if signFound {
                return 0
            }
            signFound = true
            continue
        }

        if isMinusSign(c) {
            if signFound {
                return 0
            }
            minus = true
            signFound = true
            continue
        }

        if !isDigit(c) {
            return 0
        }

        break
    }

    num, ok := strToInt(str[start:], minus)
    if !ok {
        return 0
    }
    return num
}

func strToInt(s string, minus bool) (num int, ok bool) {
    var result int64

    for _, c := range s {
        b := byte(c)
        if !isDigit(b) && !isWhiteSpace(b) {
            break
        }

        if isWhiteSpace(b) {
            break
        }

        d := c - '0'
        result = result * 10 + int64(d)

        if !minus && result > math.MaxInt32 {
            return math.MaxInt32, true
        } else if minus && -result < math.MinInt32 {
            return math.MinInt32, true
        }
    }

    if minus {
        return -int(result), true
    }
    return int(result), true
}

func isDigit(c byte) bool {
    if c - '0' >= 0 && c - '0' <= 9 {
        return true
    }
    return false
}

func isWhiteSpace(c byte) bool {
    if c == ' ' {
        return true
    }
    return false
}

func isPlusSign(c byte) bool {
    if c == '+' {
        return true
    }
    return false
}

func isMinusSign(c byte) bool {
    if c == '-' {
        return true
    }
    return false
}
