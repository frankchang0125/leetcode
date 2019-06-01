/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 *
 * https://leetcode.com/problems/reverse-string-ii/description/
 *
 * algorithms
 * Easy (45.43%)
 * Likes:    243
 * Dislikes: 734
 * Total Accepted:    59.4K
 * Total Submissions: 130.7K
 * Testcase Example:  '"abcdefg"\n2'
 *
 *
 * Given a string and an integer k, you need to reverse the first k characters
 * for every 2k characters counting from the start of the string. If there are
 * less than k characters left, reverse all of them. If there are less than 2k
 * but greater than or equal to k characters, then reverse the first k
 * characters and left the other as original.
 *
 *
 * Example:
 *
 * Input: s = "abcdefg", k = 2
 * Output: "bacdfeg"
 * 
 * 
 * 
 * Restrictions:
 *
 * ⁠The string consists of lower English letters only.
 * ⁠Length of the given string and k will in the range [1, 10000]
 * 
 */
func reverseStr(s string, k int) string {
    S := []byte(s)
    N := len(s)
    if N == 1 {
        return s
    }

    i := 0
    j := 0

    for j < N {
        if j + 2 * k <= N || j + k <= N {
            j = j + 2 * k
            reverse(S, i, i + k - 1)
        } else {
            j = N
            reverse(S, i, j - 1)
        }

        i = j
    }

    return string(S)
}

func reverse(s []byte, start int, end int) {
    N := len(s)
    if N < 2 {
        return
    } else if end >= N {
        return
    }

    l := start
    r := end

    for l < r {
        s[l], s[r] = s[r], s[l]
        l++
        r--
    }
}
