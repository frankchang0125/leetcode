/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 *
 * https://leetcode.com/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (47.88%)
 * Likes:    2349
 * Dislikes: 257
 * Total Accepted:    183.8K
 * Total Submissions: 372.7K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * Given an array nums containing n + 1 integers where each integer is between
 * 1 and n (inclusive), prove that at least one duplicate number must exist.
 * Assume that there is only one duplicate number, find the duplicate one.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,3,4,2,2]
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,1,3,4,2]
 * Output: 3
 * 
 * Note:
 * 
 * 
 * You must not modify the array (assume the array is read only).
 * You must use only constant, O(1) extra space.
 * Your runtime complexity should be less than O(n^2).
 * There is only one duplicate number in the array, but it could be repeated
 * more than once.
 * 
 * 
 */
// If no duplicate numbers in the array,
// then numbers within range: 1->n will have only n numbers.
// Since there're n+1 numbers, duplicated number must exists.
//
// Let m = Middle number of numbers range: 1->n
// We can divide the numbers: 1->n into two parts: (1->m) & (m+1->n)
//
// If the duplicate number is in range: 1->m,
// then the count of numbers in array which are less or equal to m
// will greater than m.
//  ex: [1, 2, 2, 3, 4], n = 4, divide into: (1->2) & (3->4), m = 2
//      Count of numbers less or equal to 2: 3 (1, 2, 2), which > m,
//      the duplicate number must in left part, i.e. range: 1->2
//
// If the duplicate number is in range: m+1->n,
// then the count of numbers in array which are less or equal to m
// will equal or less than m.
//  ex: [1, 2, 3, 3, 4], n = 4, divide into: (1->2) & (3->4), m = 2
//      Count of numbers less or equal to 2: 2 (1, 2), which <= m,
//      the duplicate number must in right part, i.e. range: 3->4
//
//  ex: arr = [2, 3, 5, 4, 3, 2, 6, 7], numbers range: 1->7
//      Middle number of 1->7: m = 4, i.e. (1->4) & (5->7)
//      Count of numbers less or equal to 4: 5 (2, 3, 4, 3, 2), which > m
//      The duplicate number must in left part, i.e. range: 1->4
//      l = 1, r = 4
//
//      Middle number of 1->4: m = 2, i.e. (1->2) & (3->4)
//      Count of numbers less or equal to 2: 2 (2, 2), which = m
//      The duplicate number must in right part, i.e. range: 3->4
//      l = 3, r = 4
//
//      Middle number of 3->4: m = 3, i.e. (3) & (4)
//      Count of numbers less or equal to 3: 4 (2, 3, 3, 2), which > m
//      The duplicate number must in left part, i.e. range: 3
//      l = 3, r = 3
//
//      Answer = 3
//
// Note: This method cannot guarantee to find all duplicate numbers.
//       ex: [2, 3, 5, 4, 3, 2, 6, 7], for range: (1->2) & (3->4), m = 2
//           Count of numbers less or equal to 2: 2 (2, 2), which = m.
//           However, we cannot distinguish whether its: [1, 2] or [2, 2],
//           since count <= 2 of both cases are both 2.
//
// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func findDuplicate(nums []int) int {
    N := len(nums)
    if N < 1 {
        return -1 
    }

    l := 1
    r := N - 1

    for l < r {
        // Divide number into: (l->m) & (m+1->r)
        mid := l + (r - l) / 2

        // Count numbers which is less or equal to m
        count := 0
        for _, num := range nums {
            if num <= mid {
                count++
            }
        }

        if count > mid {
            // Duplicate number must in left part, i.e. range: l->m
            r = mid
        } else {
            // Duplicate number must in right part, i.e. range: m+1->r
            l = mid + 1
        }
    }
    
    return l
}
