/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 *
 * https://leetcode.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (22.54%)
 * Total Accepted:    227.2K
 * Total Submissions: 1M
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * Given two words (beginWord and endWord), and a dictionary's word list, find
 * the length of shortest transformation sequence from beginWord to endWord,
 * such that:
 *
 *
 * Only one letter can be changed at a time.
 * Each transformed word must exist in the word list. Note that beginWord is
 * not a transformed word.
 *
 *
 * Note:
 *
 *
 * Return 0 if there is no such transformation sequence.
 * All words have the same length.
 * All words contain only lowercase alphabetic characters.
 * You may assume no duplicates in the word list.
 * You may assume beginWord and endWord are non-empty and are not the same.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * Output: 5
 *
 * Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" ->
 * "dog" -> "cog",
 * return its length 5.
 *
 *
 * Example 2:
 *
 *
 * Input:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * Output: 0
 *
 * Explanation: The endWord "cog" is not in wordList, therefore no possible
 * transformation.
 *
 *
 *
 *
 *
 */
// Reference: https://goo.gl/ZtPLUz
func ladderLength(beginWord string, endWord string, wordList []string) int {
	return bidirectionalBFS(beginWord, endWord, wordList)
}

// Expands the words from both 'beginWord set' and 'endWord set'
// in different directions. Try to find the transformed word
// which is in 'both' two directional sets
func bidirectionalBFS(beginWord string, endWord string, wordList []string) int {
	wordDict := make(map[string]bool)
	// Preprocesssing, create word list dictionary.
	for _, word := range wordList {
		wordDict[word] = true
	}

	// End word doesn't even is in word list, just return.
	if _, ok := wordDict[endWord]; !ok {
		return 0
	}

	s1 := make(map[string]bool)
	s1[beginWord] = true
	delete(wordDict, beginWord)

	s2 := make(map[string]bool)
	s2[endWord] = true
	delete(wordDict, endWord)

	current := s1
	other := s2
	steps := 0

	for {
		if len(current) == 0 || len(other) == 0 {
			return 0
		}

		steps++

		// Always expand the smaller set first.
		if len(current) > len(other) {
			tmp := current
			current = other
			other = tmp
		}

		// The valid transformed words set during this expansion.
		s := make(map[string]bool)

		for word := range current {
			w := []byte(word)

			for i := 0; i < len(w); i++ {
				// Transform character from 'a' to 'z'.
				for j := 0; j < 26; j++ {
					w[i] = byte(j) + 'a'
					wStr := string(w)

					if _, ok := other[wStr]; ok {
						// Transformed word is in other set,
						// increase steps, which means we can find
						// the word in other set next step.
						return steps + 1
					}

					if _, ok := wordDict[wStr]; ok {
						s[wStr] = true

						// Each word in the dictionary should be used once only,
						// thus delete it from dictionary.
						delete(wordDict, wStr)
					}
				}

				// Restore the transformed word back.
				w[i] = word[i]
			}
		}

		// Replace current set with the new valid transformed words set.
		current = s
	}
}
