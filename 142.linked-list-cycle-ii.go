/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 *
 * https://leetcode.com/problems/linked-list-cycle-ii/description/
 *
 * algorithms
 * Medium (31.92%)
 * Likes:    1381
 * Dislikes: 99
 * Total Accepted:    212.4K
 * Total Submissions: 665K
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * Given a linked list, return the node where the cycle begins. If there is no
 * cycle, return null.
 * 
 * To represent a cycle in the given linked list, we use an integer pos which
 * represents the position (0-indexed) in the linked list where tail connects
 * to. If pos is -1, then there is no cycle in the linked list.
 * 
 * Note: Do not modify the linked list.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: head = [3,2,0,-4], pos = 1
 * Output: tail connects to node index 1
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * second node.
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: head = [1,2], pos = 0
 * Output: tail connects to node index 0
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * first node.
 * 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: head = [1], pos = -1
 * Output: no cycle
 * Explanation: There is no cycle in the linked list.
 * 
 * 
 * 
 * 
 * 
 * 
 * Follow up:
 * Can you solve it without using extra space?
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 1. Use fast, slow pointers to determine whether there's a cycle in the linked list.
//    Fast pointer moves 2 steps a time, slow pointer moves 1 step a time.
//    If fast pointer catch up slow pointer, then there's a cycle in the linked list
//    and the match up node of fast pointer and slow pointer must in the cycle.
//    Return the match up node if cycle exists.
// 2. Start from the match up node, keep iterating until the match up node is visited again.
//    During the iteration, count the number of nodes: 'n' in the cycle.
// 3. Use another set of two pointers, let first pointer moves 'n' steps first.
//    Then moves second pointer along with first pointers. When these two pointers meet.
//    It will be the entry point of the cycle.
//      => Total 'k' nodes, first pointer moves 'n' steps first, then both pointers will meet at:
//         (k - n)th node, which is also the entry node of the cycle.
func detectCycle(head *ListNode) *ListNode {
    nodeInCycle := hasCycle(head)
    if nodeInCycle == nil {
        return nil
    }

    n := numNodesInCycle(nodeInCycle)

    p1 := head
    p2 := head

    for i := 0; i < n; i++ {
        p1 = p1.Next
    }

    for p1 != p2 {
        p1 = p1.Next
        p2 = p2.Next
    }

    return p1
}

func hasCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    slow := head.Next
    if slow == nil {
        return nil
    }

    fast := slow.Next

    for fast != nil && slow != nil {
        if fast == slow {
            return fast
        }

        slow = slow.Next
        
        fast = fast.Next
        if fast != nil {
            fast = fast.Next
        }
    }

    return nil
}

func numNodesInCycle(node *ListNode) int {
    cur := node.Next
    count := 1

    for cur != node {
        count++
        cur = cur.Next
    }

    return count
}
