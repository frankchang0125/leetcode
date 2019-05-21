/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 *
 * https://leetcode.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (49.93%)
 * Likes:    971
 * Dislikes: 76
 * Total Accepted:    264.5K
 * Total Submissions: 529.5K
 * Testcase Example:  '"leetcode"'
 *
 * 
 * Given a string, find the first non-repeating character in it and return it's
 * index. If it doesn't exist, return -1.
 * 
 * Examples:
 * 
 * s = "leetcode"
 * return 0.
 * 
 * s = "loveleetcode",
 * return 2.
 * 
 * 
 * 
 * 
 * Note: You may assume the string contain only lowercase letters.
 * 
 */
func firstUniqChar(s string) int {
    m := map[rune]int{}

    for _, c := range s {
        if _, ok := m[c]; !ok {
            m[c] = 0
        }
        m[c]++
    }

    for i, c := range s {
        if num, ok := m[c]; ok {
            if num == 1 {
                return i
            }
        }
    }

    return -1
}
