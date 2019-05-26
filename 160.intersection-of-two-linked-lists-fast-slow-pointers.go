/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 *
 * https://leetcode.com/problems/intersection-of-two-linked-lists/description/
 *
 * algorithms
 * Easy (33.67%)
 * Likes:    2070
 * Dislikes: 175
 * Total Accepted:    302.7K
 * Total Submissions: 896.4K
 * Testcase Example:  '8\n[4,1,8,4,5]\n[5,0,1,8,4,5]\n2\n3'
 *
 * Write a program to find the node at which the intersection of two singly
 * linked lists begins.
 * 
 * For example, the following two linked lists:
 * 
 * 
 * begin to intersect at node c1.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * 
 * Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA =
 * 2, skipB = 3
 * Output: Reference of the node with value = 8
 * Input Explanation: The intersected node's value is 8 (note that this must
 * not be 0 if the two lists intersect). From the head of A, it reads as
 * [4,1,8,4,5]. From the head of B, it reads as [5,0,1,8,4,5]. There are 2
 * nodes before the intersected node in A; There are 3 nodes before the
 * intersected node in B.
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * 
 * Input: intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3,
 * skipB = 1
 * Output: Reference of the node with value = 2
 * Input Explanation: The intersected node's value is 2 (note that this must
 * not be 0 if the two lists intersect). From the head of A, it reads as
 * [0,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes
 * before the intersected node in A; There are 1 node before the intersected
 * node in B.
 * 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * 
 * Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB =
 * 2
 * Output: null
 * Input Explanation: From the head of A, it reads as [2,6,4]. From the head of
 * B, it reads as [1,5]. Since the two lists do not intersect, intersectVal
 * must be 0, while skipA and skipB can be arbitrary values.
 * Explanation: The two lists do not intersect, so return null.
 * 
 * 
 * 
 * 
 * Notes:
 * 
 * 
 * If the two linked lists have no intersection at all, return null.
 * The linked lists must retain their original structure after the function
 * returns.
 * You may assume there are no cycles anywhere in the entire linked
 * structure.
 * Your code should preferably run in O(n) time and use only O(1) memory.
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
// Calculate the length of two linked lists.
// For longer lined list, walk 'k' steps from head first,
// where k is the length difference of these two linked list.
// Then walk longer and shorter linked list together, they will reach to the
// end of lists at same time.
// We can find the intersection node during the walk.
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    // Calculate the length of two linked list.
    length1 := listLength(headA)
    length2 := listLength(headB)

    // Calculate the length difference: 'k' of these tow linked list.
    lengthDiff := abs(length1 - length2)

    var fast *ListNode
    var slow *ListNode

    if length1 >= length2 {
        fast = headA
        slow = headB
    } else {
        fast = headB
        slow = headA
    }

    // Walk longer linked list 'k' steps first.
    for i := 0; i < lengthDiff; i++ {
        fast = fast.Next
    }

    // Find longer and shower linked lists together to find the intersection node.
    for fast != nil && slow != nil {
        if fast == slow {
            return fast
        }

        fast = fast.Next
        slow = slow.Next
    }

    return nil
}

func listLength(head *ListNode) int {
    cur := head
    length := 0

    for cur != nil {
        length++
        cur = cur.Next
    }
    return length
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
