/*
 * @lc app=leetcode id=845 lang=golang
 *
 * [845] Longest Mountain in Array
 *
 * https://leetcode.com/problems/longest-mountain-in-array/description/
 *
 * algorithms
 * Medium (34.16%)
 * Likes:    313
 * Dislikes: 13
 * Total Accepted:    18K
 * Total Submissions: 52.4K
 * Testcase Example:  '[2,1,4,7,3,2,5]'
 *
 * Let's call any (contiguous) subarray B (of A) a mountain if the following
 * properties hold:
 * 
 * 
 * B.length >= 3
 * There exists some 0 < i < B.length - 1 such that B[0] < B[1] < ... B[i-1] <
 * B[i] > B[i+1] > ... > B[B.length - 1]
 * 
 * 
 * (Note that B could be any subarray of A, including the entire array A.)
 * 
 * Given an array A of integers, return the length of the longest mountain. 
 * 
 * Return 0 if there is no mountain.
 * 
 * Example 1:
 * 
 * 
 * Input: [2,1,4,7,3,2,5]
 * Output: 5
 * Explanation: The largest mountain is [1,4,7,3,2] which has length 5.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [2,2,2]
 * Output: 0
 * Explanation: There is no mountain.
 * 
 * 
 * Note:
 * 
 * 
 * 0 <= A.length <= 10000
 * 0 <= A[i] <= 10000
 * 
 * 
 * Follow up:
 * 
 * 
 * Can you solve it using only one pass?
 * Can you solve it in O(1) space?
 * 
 * 
 */
// References:
//  http://bit.ly/2InWflg
//  http://bit.ly/2InjQ5C
func longestMountain(A []int) int {
    N := len(A)
    if N < 3 {
        return 0
    }

    inc := 0
    dec := 0
    ans := 0

    for i := 1; i < N; i++ {
        if dec > 0 && A[i] > A[i-1] || A[i] == A[i-1] {
            inc = 0
            dec = 0
        }
        if A[i] > A[i-1] {
            inc += 1
        }
        if A[i] < A[i-1] {
            dec += 1
        }

        if inc > 0 && dec > 0 {
            ans = max(ans, inc + dec + 1)
        }
    }

    if ans < 3 {
        return 0
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
