/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 *
 * https://leetcode.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (30.66%)
 * Likes:    2324
 * Dislikes: 157
 * Total Accepted:    242.6K
 * Total Submissions: 782.5K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * Given a string S and a string T, find the minimum window in S which will
 * contain all the characters in T in complexity O(n).
 * 
 * Example:
 * 
 * 
 * Input: S = "ADOBECODEBANC", T = "ABC"
 * Output: "BANC"
 * 
 * 
 * Note:
 * 
 * 
 * If there is no such window in S that covers all characters in T, return the
 * empty string "".
 * If there is such window, you are guaranteed that there will always be only
 * one unique minimum window in S.
 * 
 * 
 */
// References:
//  http://bit.ly/2FCMiyK
//  http://bit.ly/2FA24dx
//  http://bit.ly/2FBpkYU
func minWindow(s string, t string) string {
    N := len(s)
    m := map[byte]int{}
    needs := map[byte]int{}
    match := 0

    for i := 0; i < len(t); i++ {
        needs[t[i]]++
    }

    i := 0
    j := 0
    minLen := math.MaxInt32
    var minSubString string

    for ; j < N; j++ {
        c := s[j]
        if _, ok := needs[c]; ok {
            m[c]++
            if m[c] == needs[c] {
                match++
            }
        }

        for match == len(needs) {
            c = s[i]
            if _, ok := needs[c]; ok {
                if m[c] == needs[c] {
                    strLen := j - i + 1
                    if strLen < minLen {
                        minSubString = s[i:j+1]
                        minLen = strLen
                    }
                    match--
                }
            }

            m[c]--
            i++
        }
    }

    return minSubString
}
