/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 *
 * https://leetcode.com/problems/odd-even-linked-list/description/
 *
 * algorithms
 * Medium (49.01%)
 * Likes:    716
 * Dislikes: 219
 * Total Accepted:    145.3K
 * Total Submissions: 296.4K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Given a singly linked list, group all odd nodes together followed by the
 * even nodes. Please note here we are talking about the node number and not
 * the value in the nodes.
 * 
 * You should try to do it in place. The program should run in O(1) space
 * complexity and O(nodes) time complexity.
 * 
 * Example 1:
 * 
 * 
 * Input: 1->2->3->4->5->NULL
 * Output: 1->3->5->2->4->NULL
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 2->1->3->5->6->4->7->NULL
 * Output: 2->3->6->7->1->5->4->NULL
 * 
 * 
 * Note:
 * 
 * 
 * The relative order inside both the even and odd groups should remain as it
 * was in the input.
 * The first node is considered odd, the second node even and so on ...
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
// References:
//  http://bit.ly/2Me8dlo
//  http://bit.ly/2w8B2Ff
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    // odd always point to the last processed odd node.
    // even always point to the last processed event node.
    odd := head
    even := head.Next

    for even != nil && even.Next != nil {
        tmp := odd.Next
        odd.Next = even.Next
        even.Next = even.Next.Next
        odd.Next.Next = tmp
        odd = odd.Next
        even = even.Next
    }

    return head
}
