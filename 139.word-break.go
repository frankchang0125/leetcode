/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 *
 * https://leetcode.com/problems/word-break/description/
 *
 * algorithms
 * Medium (33.99%)
 * Total Accepted:    303K
 * Total Submissions: 882.9K
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * Given a non-empty string s and a dictionary wordDict containing a list of
 * non-empty words, determine if s can be segmented into a space-separated
 * sequence of one or more dictionary words.
 *
 * Note:
 *
 *
 * The same word in the dictionary may be reused multiple times in the
 * segmentation.
 * You may assume the dictionary does not contain duplicate words.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "leetcode", wordDict = ["leet", "code"]
 * Output: true
 * Explanation: Return true because "leetcode" can be segmented as "leet
 * code".
 *
 *
 * Example 2:
 *
 *
 * Input: s = "applepenapple", wordDict = ["apple", "pen"]
 * Output: true
 * Explanation: Return true because "applepenapple" can be segmented as "apple
 * pen apple".
 * Note that you are allowed to reuse a dictionary word.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * Output: false
 *
 *
 */
type Solution struct {
	wordDictMap map[string]bool
	mem         map[string]bool // Memorize previous word break results.
}

// Reference: https://goo.gl/4PnXAS
func wordBreak(s string, wordDict []string) bool {
	solution := Solution{
		wordDictMap: make(map[string]bool),
		mem:         make(map[string]bool),
	}

	// Preprocessing, build word dictionary map.
	for _, word := range wordDict {
		solution.wordDictMap[word] = true
	}

	return solution.findWordBreak(s)
}

func (s Solution) findWordBreak(w string) bool {
	// Try to find the word break result from memoization.
	if m, ok := s.mem[w]; ok {
		return m
	}

	// If the word break can be found in word dictionary, return true.
	if _, found := s.wordDictMap[w]; found {
		return true
	}

	var result bool

	// Try to break the word one-by-one from left to right,
	// return true if both:
	// 	1. Left part of the break word is in dictionary.
	//	2. Right part of the break word or one of its break-word-subsets
	//	   is in dictionary. (Right part may be break into more sub-words...)
	// (Note: Left and right can be exchanged.)
	for i := 1; i < len(w); i++ {
		result = s.inDict(w[:i]) && s.findWordBreak(w[i:])
		if result {
			break
		}
	}

	// Memorize the word break result.
	s.mem[w] = result
	return result
}

func (s Solution) inDict(w string) bool {
	_, found := s.wordDictMap[w]
	return found
}
