/*
 * @lc app=leetcode id=676 lang=golang
 *
 * [676] Implement Magic Dictionary
 *
 * https://leetcode.com/problems/implement-magic-dictionary/description/
 *
 * algorithms
 * Medium (50.84%)
 * Total Accepted:    22.3K
 * Total Submissions: 43.9K
 * Testcase Example:  '["MagicDictionary", "buildDict", "search", "search", "search", "search"]\n[[], [["hello","leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]'
 *
 *
 * Implement a magic directory with buildDict, and search methods.
 *
 *
 *
 * For the method buildDict, you'll be given a list of non-repetitive words to
 * build a dictionary.
 *
 *
 *
 * For the method search, you'll be given a word, and judge whether if you
 * modify exactly one character into another character in this word, the
 * modified word is in the dictionary you just built.
 *
 *
 * Example 1:
 *
 * Input: buildDict(["hello", "leetcode"]), Output: Null
 * Input: search("hello"), Output: False
 * Input: search("hhllo"), Output: True
 * Input: search("hell"), Output: False
 * Input: search("leetcoded"), Output: False
 *
 *
 *
 * Note:
 *
 * You may assume that all the inputs are consist of lowercase letters a-z.
 * For contest purpose, the test data is rather small by now. You could think
 * about highly efficient algorithm after the contest.
 * Please remember to RESET your class variables declared in class
 * MagicDictionary, as static/class variables are persisted across multiple
 * test cases. Please see here for more details.
 *
 *
 */
// Build magic dictionary with fuzzy character: '*' of each character
// of the words in magic dictionary.
// ex: 'hello' -> '*ello', 'h*llo', 'he*lo', 'hel*o', 'hell*'.
//
// Also, since we must modify exactly one character into another character
// in the searched word, thus, we also have to memorize the original
// characters when building map with fuzzy character: '*'.
//
// When searching, replace each character with fuzzy character: '*' in
// searched word and try to look up them in the magic dictionary.
// If found and the replaced character is not the only character in original
// characters list, return true, otherwise, return false.
//
// Reference: https://goo.gl/UXmWCV
type MagicDictionary struct {
	m map[string]map[byte]bool
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{
		m: make(map[string]map[byte]bool),
	}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for _, word := range dict {
		w := []byte(word)

		// Build magic dictionary by replacing each character
		// with fuzzy character: '*'.
		for i := range word {
			w[i] = '*'

			fuzzy, ok := this.m[string(w)]
			if !ok {
				fuzzy = make(map[byte]bool)
				this.m[string(w)] = fuzzy
			}

			// Memorize the original character.
			fuzzy[word[i]] = true

			// Restore original character for next interation.
			w[i] = word[i]
		}
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	b := []byte(word)

	// Replace each character with fuzzy character: '*'
	// and look up in magic dictionary.
	for i, _ := range word {
		b[i] = '*'

		if fuzzy, ok := this.m[string(b)]; ok {
			// More the one original characters of the fuzzy word exists,
			// the word can be found in the magic dictionary.
			if len(fuzzy) > 1 {
				return true
			}

			// The replaced character is not in the original characters list
			// of the corresponded fuzzy word. The word can be found in the
			// magic dictionary.
			if _, found := fuzzy[word[i]]; !found {
				return true
			}
		}

		b[i] = word[i]
	}

	// The word cannot be found in the magic dictionary.
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
