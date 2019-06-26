/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 *
 * https://leetcode.com/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Easy (37.19%)
 * Likes:    1613
 * Dislikes: 124
 * Total Accepted:    127K
 * Total Submissions: 339.5K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * Given a string s and a non-empty string p, find all the start indices of p's
 * anagrams in s.
 * 
 * Strings consists of lowercase English letters only and the length of both
 * strings s and p will not be larger than 20,100.
 * 
 * The order of output does not matter.
 * 
 * Example 1:
 * 
 * Input:
 * s: "cbaebabacd" p: "abc"
 * 
 * Output:
 * [0, 6]
 * 
 * Explanation:
 * The substring with start index = 0 is "cba", which is an anagram of "abc".
 * The substring with start index = 6 is "bac", which is an anagram of
 * "abc".
 * 
 * 
 * 
 * Example 2:
 * 
 * Input:
 * s: "abab" p: "ab"
 * 
 * Output:
 * [0, 1, 2]
 * 
 * Explanation:
 * The substring with start index = 0 is "ab", which is an anagram of "ab".
 * The substring with start index = 1 is "ba", which is an anagram of "ab".
 * The substring with start index = 2 is "ab", which is an anagram of "ab".
 * 
 * 
 */
func findAnagrams(s string, p string) []int {
    N := len(s)
    m := map[byte]int{}
    needs := map[byte]int{}
    match := 0

    for i := 0; i < len(p); i++ {
        needs[p[i]]++
    }

    i := 0
    j := 0
    ans := []int{}

    for ; j < N; j++ {
        c := s[j]
        if _, ok := needs[c]; ok {
            m[c]++
            if m[c] == needs[c] {
                match++
            }
        }

        for match == len(needs) {
            if j - i + 1 == len(p) {
                ans = append(ans, i)
            }

            c = s[i]
            if _, ok := needs[c]; ok {
                if m[c] == needs[c] {
                    match--
                }
                m[c]--
            }

            i++
        }
    }

    return ans
}
