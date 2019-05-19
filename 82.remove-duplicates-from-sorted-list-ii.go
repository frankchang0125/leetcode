/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (32.92%)
 * Likes:    801
 * Dislikes: 71
 * Total Accepted:    180.3K
 * Total Submissions: 547.6K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * Given a sorted linked list, delete all nodes that have duplicate numbers,
 * leaving only distinct numbers from the original list.
 * 
 * Example 1:
 * 
 * 
 * Input: 1->2->3->3->4->4->5
 * Output: 1->2->5
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 1->1->1->2->3
 * Output: 2->3
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
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    cur := head
    var prev *ListNode

    for cur != nil {
        if cur.Next != nil && cur.Next.Val == cur.Val {
            // Duplicate node found, find the end node of the duplicate nodes.
            for cur.Next != nil && cur.Val == cur.Next.Val {
                cur = cur.Next
            }

            if prev != nil {
                // Remove the duplicate nodes, prev node remain the same.
                prev.Next = cur.Next
            } else {
                // prev is nil, which mean head node is the duplicate node,
                // update head to the next node.
                head = cur.Next
            }
        } else {
            prev = cur
        }

        cur = cur.Next
    }

    return head
}
