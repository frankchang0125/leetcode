/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 *
 * https://leetcode.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (27.04%)
 * Likes:    592
 * Dislikes: 809
 * Total Accepted:    191.5K
 * Total Submissions: 706.5K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, rotate the list to the right by k places, where k is
 * non-negative.
 * 
 * Example 1:
 * 
 * 
 * Input: 1->2->3->4->5->NULL, k = 2
 * Output: 4->5->1->2->3->NULL
 * Explanation:
 * rotate 1 steps to the right: 5->1->2->3->4->NULL
 * rotate 2 steps to the right: 4->5->1->2->3->NULL
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 0->1->2->NULL, k = 4
 * Output: 2->0->1->NULL
 * Explanation:
 * rotate 1 steps to the right: 2->0->1->NULL
 * rotate 2 steps to the right: 1->2->0->NULL
 * rotate 3 steps to the right: 0->1->2->NULL
 * rotate 4 steps to the right: 2->0->1->NULL
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }

    if k == 0 {
        return head
    }

    cur := head
    var oldTail *ListNode
    length := 0

    for cur != nil {
        length++
        oldTail = cur
        cur = cur.Next
    }

    steps := length - (k % length)
    if steps == length {
        // No need to rotate.
        return head
    }

    var newTail *ListNode
    cur = head

    for i := 0; i < steps; i++ {
        newTail = cur
        cur = cur.Next
    }

    if newTail != nil {
        newTail.Next = nil
    }

    if oldTail != nil {
        oldTail.Next = head
    }

    return cur
}
