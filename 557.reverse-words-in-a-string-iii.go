/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 *
 * https://leetcode.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (63.96%)
 * Likes:    615
 * Dislikes: 66
 * Total Accepted:    124.8K
 * Total Submissions: 194.9K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * Given a string, you need to reverse the order of characters in each word
 * within a sentence while still preserving whitespace and initial word order.
 * 
 * Example 1:
 * 
 * Input: "Let's take LeetCode contest"
 * Output: "s'teL ekat edoCteeL tsetnoc"
 * 
 * 
 * 
 * Note:
 * In the string, each word is separated by single space and there will not be
 * any extra space in the string.
 * 
 */
func reverseWords(s string) string {
    N := len(s)
    if N < 2 {
        return s
    }

    S := []byte(s)
    start := 0
    end := 0

    for start < N {
        for end < N && S[end] != ' ' {
            end++
        }

        reverse(S, start, end - 1)
        start = end + 1
        end++
    }

    return string(S)
}

func reverse(s []byte, start int, end int) {
    if len(s) < 2 {
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
