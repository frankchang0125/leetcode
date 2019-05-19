/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 *
 * https://leetcode.com/problems/squares-of-a-sorted-array/description/
 *
 * algorithms
 * Easy (74.33%)
 * Likes:    235
 * Dislikes: 34
 * Total Accepted:    59K
 * Total Submissions: 81.6K
 * Testcase Example:  '[-4,-1,0,3,10]'
 *
 * Given an array of integers A sorted in non-decreasing order, return an array
 * of the squares of each number, also in sorted non-decreasing order.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: [-4,-1,0,3,10]
 * Output: [0,1,9,16,100]
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [-7,-3,2,3,11]
 * Output: [4,9,9,49,121]
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= A.length <= 10000
 * -10000 <= A[i] <= 10000
 * A is sorted in non-decreasing order.
 * 
 * 
 * 
 */
func sortedSquares(A []int) []int {
    N := len(A)
    i := 0
    j := N - 1
    pos := N - 1
    ans := make([]int, N)

    for i < j {
        if abs(A[i]) > abs(A[j]) {
            ans[pos] = A[i] * A[i]
            i++
            pos--
        } else {
            ans[pos] = A[j] * A[j]
            j--
            pos--
        }
    }

    ans[pos] = A[i] * A[i]

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
