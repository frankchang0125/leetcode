/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 *
 * https://leetcode.com/problems/lru-cache/description/
 *
 * algorithms
 * Hard (23.52%)
 * Total Accepted:    277.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 *
 * Design and implement a data structure for Least Recently Used (LRU) cache.
 * It should support the following operations: get and put.
 *
 *
 *
 * get(key) - Get the value (will always be positive) of the key if the key
 * exists in the cache, otherwise return -1.
 * put(key, value) - Set or insert the value if the key is not already present.
 * When the cache reached its capacity, it should invalidate the least recently
 * used item before inserting a new item.
 *
 *
 * Follow up:
 * Could you do both operations in O(1) time complexity?
 *
 * Example:
 *
 * LRUCache cache = new LRUCache( 2 // capacity );
 *
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // returns 1
 * cache.put(3, 3);    // evicts key 2
 * cache.get(2);       // returns -1 (not found)
 * cache.put(4, 4);    // evicts key 1
 * cache.get(1);       // returns -1 (not found)
 * cache.get(3);       // returns 3
 * cache.get(4);       // returns 4
 *
 *
 */
// Reference: http://bit.ly/2UE2VCX
type LRUCache struct {
	CacheMap      map[int]*list.Element
	CacheList     *list.List
	CacheCapacity int
}

// LRU node is appended/moved to the head of list.
type CacheNode struct {
	Key  int
	Val  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		CacheMap:      map[int]*list.Element{},
		CacheList: 	   list.New(),
		CacheCapacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	element, ok := this.CacheMap[key]
	if !ok {
		// Node not found.
		return -1
	}

	// Node found, move node to the head of list.
	this.MoveNodeToHead(element)
	node := element.Value.(*CacheNode)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	element, ok := this.CacheMap[key]
	if !ok {
		// Node key not found in cache,
		// remove the outdated node (if required)
		// and insert the new node.
		if this.CacheList.Len() == this.CacheCapacity {
			this.RemoveOutdatedNode()
		}

		node := &CacheNode{Key: key, Val: value}
		element := this.AddNewNode(node)
		this.CacheMap[key] = element
	} else {
		// Node key found in cache,
		// move the node to the head of list.
		this.MoveNodeToHead(element)

		// Node value may be different, update node value.
		element.Value.(*CacheNode).Val = value
	}
}

func (this *LRUCache) RemoveOutdatedNode() {
	element := this.CacheList.Back()
	this.CacheList.Remove(element)
	node := element.Value.(*CacheNode)
	delete(this.CacheMap, node.Key)
}

func (this *LRUCache) AddNewNode(node *CacheNode) (*list.Element) {
	return this.CacheList.PushFront(node)
}

func (this *LRUCache) MoveNodeToHead(element *list.Element) {
	this.CacheList.MoveToFront(element)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
