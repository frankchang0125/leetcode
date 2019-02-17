/*
 * @lc app=leetcode id=720 lang=golang
 *
 * [720] Longest Word in Dictionary
 *
 * https://leetcode.com/problems/longest-word-in-dictionary/description/
 *
 * algorithms
 * Easy (43.52%)
 * Total Accepted:    29.8K
 * Total Submissions: 68.4K
 * Testcase Example:  '["w","wo","wor","worl","world"]'
 *
 * Given a list of strings words representing an English Dictionary, find the
 * longest word in words that can be built one character at a time by other
 * words in words.  If there is more than one possible answer, return the
 * longest word with the smallest lexicographical order.  If there is no
 * answer, return the empty string.
 *
 * Example 1:
 *
 * Input:
 * words = ["w","wo","wor","worl", "world"]
 * Output: "world"
 * Explanation:
 * The word "world" can be built one character at a time by "w", "wo", "wor",
 * and "worl".
 *
 *
 *
 * Example 2:
 *
 * Input:
 * words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
 * Output: "apple"
 * Explanation:
 * Both "apply" and "apple" can be built from other words in the dictionary.
 * However, "apple" is lexicographically smaller than "apply".
 *
 *
 *
 * Note:
 * All the strings in the input will only contain lowercase letters.
 * The length of words will be in the range [1, 1000].
 * The length of words[i] will be in the range [1, 30].
 *
 */
// Sort the strings array from length long to short.
// If the length of two strings are equal,
// return the one which is lexicographically smaller.
type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	if len(s[i]) > len(s[j]) {
		return true
	} else if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	} else {
		return false
	}
}

// Trie (Prefix tree).
type Trie struct {
	root *Node
}

type Node struct {
	val       rune
	children  []*Node
	endOfWord bool
}

func NewNode(val rune) *Node {
	return &Node{
		val:      val,
		children: make([]*Node, 26),
	}
}

func Constructor() Trie {
	return Trie{
		root: NewNode(0),
	}
}

func (this *Trie) Insert(word string) {
	current := this.root

	for _, c := range word {
		index := c - 'a'
		if current.children[index] == nil {
			current.children[index] = NewNode(c)
		}

		current = current.children[index]
	}

	current.endOfWord = true
}

func (this *Trie) HasAllPrefixes(word string) bool {
	current := this.root

	// If all prefixes (i: 0 -> len(word)-1) of word
	// can be found, return true.
	for i := 0; i < len(word)-1; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
		if !current.endOfWord {
			return false
		}
	}

	return true
}

// To find the longest word which all of its prefixes
// can be found in the words array, we can build a
// Tile (Prefix tree) for all input words and search
// tile with each word's prefixes.
// (ex: "apple"'s prefixes are: "a", "ap", "app", "appl")
//
// Reference: https://goo.gl/hhN73F
func longestWord(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// Sort the strings array from length long to short,
	// lexicographical order from small to large for
	// equal length strings.
	sort.Sort(byLength(words))

	// Build Tile (Prefix tree).
	t := Constructor()
	for _, w := range words {
		t.Insert(w)
	}

	for _, w := range words {
		// Search tile with word's prefixes,
		// if all prefixes can be found in tile, return the word.
		if t.HasAllPrefixes(w) {
			return w
		}
	}

	return ""
}
