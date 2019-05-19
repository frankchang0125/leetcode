/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 *
 * https://leetcode.com/problems/sort-array-by-parity/description/
 *
 * algorithms
 * Easy (72.61%)
 * Likes:    445
 * Dislikes: 50
 * Total Accepted:    92.7K
 * Total Submissions: 127.6K
 * Testcase Example:  '[3,1,2,4]'
 *
 * Given an array A of non-negative integers, return an array consisting of all
 * the even elements of A, followed by all the odd elements of A.
 * 
 * You may return any answer array that satisfies this condition.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: [3,1,2,4]
 * Output: [2,4,3,1]
 * The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= A.length <= 5000
 * 0 <= A[i] <= 5000
 * 
 * 
 * 
 */
func sortArrayByParity(A []int) []int {
    N := len(A)
    if N <= 1 {
        return A
    }

    p1 := 0
    p2 := N - 1

    for p1 < p2 {
        // Iterate to find the odd number from left-hand side of array.
        for p1 < N && A[p1] & 1 == 0 {
            p1++
        }

        // Iterate to find the even number from right-hand side of array.
        for p2 >= 0 && A[p2] & 1 == 1 {
            p2--
        }

        if p1 < p2 {
            A[p1], A[p2] = A[p2], A[p1]
            p1++ 
            p2--
        }
    }

    return A
}
