/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (33.90%)
 * Total Accepted:    377.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, remove the n-th node from the end of list and return
 * its head.
 * 
 * Example:
 * 
 * 
 * Given linked list: 1->2->3->4->5, and n = 2.
 * 
 * After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 * 
 * 
 * Note:
 * 
 * Given n will always be valid.
 * 
 * Follow up:
 * 
 * Could you do this in one pass?
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if n == 0 {
        return head
    }

    fast := head
    var slow, prevSlow *ListNode
    count := 0

    for fast != nil {
        if count == n {
            slow = head
        }
        
        fast = fast.Next
        if slow != nil {
            prevSlow = slow
            slow = slow.Next
        }

        count++
    }

    if count > n {
        // Nth node from end of list is not the head node.
        prevSlow.Next = slow.Next
    } else if count == n {
        // Nth node from end of list is the head node.
        head = head.Next
    }

    return head
}
