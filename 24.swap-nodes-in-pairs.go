/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (42.63%)
 * Total Accepted:    272.4K
 * Total Submissions: 638.4K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 *
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 *
 *
 *
 * Example:
 *
 *
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	var prev, pprev *ListNode
	cur := head
	i := 0

	for {
		if cur == nil {
			break
		}

		if i%2 == 1 {
			prev.Next = cur.Next
			cur.Next = prev

			if i == 1 {
				head = cur
			} else {
				pprev.Next = cur
			}

			tmp := prev
			prev = cur
			cur = tmp
		} else {
			pprev = prev
		}

		prev = cur
		cur = cur.Next
		i++
	}

	return head
}
