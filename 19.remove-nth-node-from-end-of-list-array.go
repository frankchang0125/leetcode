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

    arr := []*ListNode{}
    cur := head

    for {
        if cur == nil {
            break
        }

        arr = append(arr, cur)
        cur = cur.Next
    }

    removeNode := arr[len(arr)-n]
    
    if removeNode == head {
        head = head.Next
    } else {
        prevNode := arr[len(arr)-n-1]
        prevNode.Next = removeNode.Next
    }

    removeNode.Next = nil

    return head
}
