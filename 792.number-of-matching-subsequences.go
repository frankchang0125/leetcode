/*
 * @lc app=leetcode id=792 lang=golang
 *
 * [792] Number of Matching Subsequences
 *
 * https://leetcode.com/problems/number-of-matching-subsequences/description/
 *
 * algorithms
 * Medium (40.68%)
 * Total Accepted:    16.3K
 * Total Submissions: 39.8K
 * Testcase Example:  '"abcde"\n["a","bb","acd","ace"]'
 *
 * Given string S and a dictionary of words words, find the number of words[i]
 * that is a subsequence of S.
 * 
 * 
 * Example :
 * Input: 
 * S = "abcde"
 * words = ["a", "bb", "acd", "ace"]
 * Output: 3
 * Explanation: There are three words in words that are a subsequence of S:
 * "a", "acd", "ace".
 * 
 * 
 * Note:
 * 
 * 
 * All words in words and S will only consists of lowercase letters.
 * The length of S will be in the range of [1, 50000].
 * The length of words will be in the range of [1, 5000].
 * The length of words[i] will be in the range of [1, 50].
 * 
 * 
 */
// Reference: https://goo.gl/pcTpgL
func numMatchingSubseq(S string, words []string) int {
    pos := make(map[rune][]int)
    
    // Preprocessing, create the position index map of each character in S.
    for i, s := range S {
        p, ok := pos[s]
        if !ok {
            p = make([]int, 0)
            pos[s] = p
        }
        pos[s] = append(p, i)
    }
    
    count := 0
    for _, word := range words {
        if isMatch(pos, word) {
            count++
        }
    }
    
    return count 
}

func isMatch(pos map[rune][]int, word string) bool {
    // prevIndex records the position index of previous character found in word.
    // Therefore, we should find a position index which is larger than prevIndex in
    // pos[c] for next character: c.
    prevIndex := -1

    for _, c := range word {
        // Current character not even exists in position index map,
        // word must not a subsequence of S.
        p, ok := pos[c]
        if !ok {
            return false
        }
        
        // Try to find a position index which is larger than prevIndex in pos[c].
        found := false
        for i := 0; i < len(p); i++ {
            if p[i] > prevIndex {
                prevIndex = p[i]
                found = true
                break
            }
        }
        
        if !found {
            // Cannot found, word must not a subsequence of S.
            return false
        }
    }
    
    // All characters of word can be found in position index map, word is a
    // subsequence of S.
    return true
}
