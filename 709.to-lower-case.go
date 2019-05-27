/*
 * @lc app=leetcode id=709 lang=golang
 *
 * [709] To Lower Case
 *
 * https://leetcode.com/problems/to-lower-case/description/
 *
 * algorithms
 * Easy (76.77%)
 * Likes:    272
 * Dislikes: 904
 * Total Accepted:    105K
 * Total Submissions: 136.6K
 * Testcase Example:  '"Hello"'
 *
 * Implement function ToLowerCase() that has a string parameter str, and
 * returns the same string in lowercase.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "Hello"
 * Output: "hello"
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "here"
 * Output: "here"
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "LOVELY"
 * Output: "lovely"
 * 
 * 
 * 
 * 
 * 
 */
func toLowerCase(str string) string {
    s := []byte(str)
    for i := 0; i < len(s); i++ {
        s[i] = lower(s[i])
    }
    return string(s)
}

func lower(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        return c - 'A' + 'a'
    }
    return c
}
