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
	return bfs(beginWord, endWord, wordList)
}

func bfs(beginWord string, endWord string, wordList []string) int {
	wordDict := make(map[string]bool)
	// Preprocesssing, create word list dictionary.
	for _, word := range wordList {
		wordDict[word] = true
	}

	// End word doesn't even is in word list, just return.
	if _, ok := wordDict[endWord]; !ok {
		return 0
	}

	queue := []string{}
	queue = append(queue, beginWord) // Enqueue
	steps := 0

	for {
		if len(queue) == 0 {
			return 0
		}

		// steps is the number of steps so far has been transformed.
		steps++
		for i := len(queue); i > 0; i-- {
			word := queue[0]
			queue = queue[1:] // Dequeue

			w := []byte(word)

			// Empty the queue of same level.
			// (i.e. All words with same index of character been transformed.)
			for j := 0; j < len(w); j++ {
				// Transform character from 'a' to 'z'.
				for k := 0; k < 26; k++ {
					w[j] = byte(k) + 'a'
					wStr := string(w)
					if wStr == endWord {
						// Transformed word is the end word,
						// increase steps, which means we can find end word
						// in the next step.
						return steps + 1
					}

					_, ok := wordDict[wStr]
					if ok && wStr != word {
						// Each word in the dictionary should be used once only,
						// thus delete it from dictionary.
						delete(wordDict, wStr)
						queue = append(queue, wStr) // Enqueue
					}
				}

				// Restore the transformed word back.
				w[j] = word[j]
			}
		}
	}
}
