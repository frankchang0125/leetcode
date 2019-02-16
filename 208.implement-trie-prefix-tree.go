/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 *
 * https://leetcode.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (36.22%)
 * Total Accepted:    159.4K
 * Total Submissions: 437.6K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * Implement a trie with insert, search, and startsWith methods.
 *
 * Example:
 *
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // returns true
 * trie.search("app");     // returns false
 * trie.startsWith("app"); // returns true
 * trie.insert("app");
 * trie.search("app");     // returns true
 *
 *
 * Note:
 *
 *
 * You may assume that all inputs are consist of lowercase letters a-z.
 * All inputs are guaranteed to be non-empty strings.
 *
 *
 */
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

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: NewNode(0),
	}
}

/** Inserts a word into the trie. */
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

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.findNode(word)
	if node == nil {
		return false
	}
	return node.endOfWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.findNode(prefix)
	return node != nil
}

func (this *Trie) findNode(word string) *Node {
	current := this.root

	for _, c := range word {
		index := c - 'a'
		if current.children[index] == nil {
			return nil
		}
		current = current.children[index]
	}

	return current
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
