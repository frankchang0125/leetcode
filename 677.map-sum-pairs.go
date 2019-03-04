/*
 * @lc app=leetcode id=677 lang=golang
 *
 * [677] Map Sum Pairs
 *
 * https://leetcode.com/problems/map-sum-pairs/description/
 *
 * algorithms
 * Medium (51.32%)
 * Total Accepted:    22.7K
 * Total Submissions: 44.3K
 * Testcase Example:  '["MapSum", "insert", "sum", "insert", "sum"]\n[[], ["apple",3], ["ap"], ["app",2], ["ap"]]'
 *
 *
 * Implement a MapSum class with insert, and sum methods.
 *
 *
 *
 * For the method insert, you'll be given a pair of (string, integer). The
 * string represents the key and the integer represents the value. If the key
 * already existed, then the original key-value pair will be overridden to the
 * new one.
 *
 *
 *
 * For the method sum, you'll be given a string representing the prefix, and
 * you need to return the sum of all the pairs' value whose key starts with the
 * prefix.
 *
 *
 * Example 1:
 *
 * Input: insert("apple", 3), Output: Null
 * Input: sum("ap"), Output: 3
 * Input: insert("app", 2), Output: Null
 * Input: sum("ap"), Output: 5
 *
 *
 *
 */
// Reference: https://goo.gl/S24t8G
type Node struct {
	children []*Node
	val      int
}

type Tile struct {
	root *Node
}

func NewTile() *Tile {
	return &Tile{
		root: &Node{
			children: make([]*Node, 128),
		},
	}
}

func NewNode() *Node {
	return &Node{
		children: make([]*Node, 128),
	}
}

func (t *Tile) Insert(s string, val int) {
	current := t.root

	for _, c := range s {
		if current.children[c] == nil {
			current.children[c] = NewNode()
		}
		current.children[c].val += val
		current = current.children[c]
	}
}

func (t *Tile) Search(s string) int {
	current := t.root

	for _, c := range s {
		if current.children[c] == nil {
			return 0
		}
		current = current.children[c]
	}

	return current.val
}

// Build tile (prefix tree) and add the value to each node
// during insertion. Search the tile to get the sum of all
// pairs' value.
type MapSum struct {
	values map[string]int
	tile   *Tile
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		values: make(map[string]int),
		tile:   NewTile(),
	}
}

func (this *MapSum) Insert(key string, val int) {
	// Check whether the given key-value pair was inserted before.
	v, ok := this.values[key]
	this.values[key] = val

	var newValue int
	if !ok {
		// Key-value pair was not inserted before, assign the value
		// to the tile node.
		newValue = val
	} else {
		// Key-value pair was inserted before, we have to substract the value
		// of original key-value pair.
		newValue = val - v
	}

	this.tile.Insert(key, newValue)
}

func (this *MapSum) Sum(prefix string) int {
	return this.tile.Search(prefix)
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
