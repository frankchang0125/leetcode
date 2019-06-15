/*
 * @lc app=leetcode id=718 lang=golang
 *
 * [718] Maximum Length of Repeated Subarray
 *
 * https://leetcode.com/problems/maximum-length-of-repeated-subarray/description/
 *
 * algorithms
 * Medium (45.86%)
 * Likes:    641
 * Dislikes: 29
 * Total Accepted:    34.3K
 * Total Submissions: 74.3K
 * Testcase Example:  '[1,2,3,2,1]\n[3,2,1,4,7]'
 *
 * Given two integer arrays A and B, return the maximum length of an subarray
 * that appears in both arrays.
 * 
 * Example 1:
 * 
 * 
 * Input:
 * A: [1,2,3,2,1]
 * B: [3,2,1,4,7]
 * Output: 3
 * Explanation: 
 * The repeated subarray with maximum length is [3, 2, 1].
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= len(A), len(B) <= 1000
 * 0 <= A[i], B[i] < 100
 * 
 * 
 * 
 * 
 */
// This is "Longest Common Substring" problem.
func findLength(A []int, B []int) int {
    N := len(A)
    M := len(B)

    // dp[i][j] = Length of longest common string of A[0->i-1] and B[0->i-1].
    // dp[i][j]:
    //  1. If i == 0 or j == 0:
    //      => Empty substring, common substring count = 0.
    //  2. If A[i-1] == B[i-1]:
    //      => dp[i-1][j-1] + 1
    //  3. If A[i-1] != B[i-1]:
    //      => Since we are looking for substring (not subsequence),
    //         reset common substring count to 0.
    //      => 0
    //  ex: A = [1,2,3], B = [2,2,3]:
    //      => dp[0][0] = 0
    //      => dp[0][1] = 0
    //      => dp[0][2] = 0
    //      => dp[0][3] = 0
    //      => dp[1][0] = 0
    //      => dp[1][1]: A[0] != B[0] = 0
    //      => dp[1][2]: A[0] != B[1] = 0
    //      => dp[1][3]: A[0] != B[2] = 0
    //      => dp[2][0] = 0
    //      => dp[2][1]: A[1] != B[0] = dp[1][0] + 1 = 1
    //      => dp[2][2]: A[1] != B[1] = dp[1][1] + 1 = 1
    //      => dp[2][3]: A[1] != B[2] = 0
    //      => dp[3][0] = 0
    //      => dp[3][1]: A[2] != B[0] = 0
    //      => dp[3][2]: A[2] != B[1] = 0
    //      => dp[3][3]: A[2] == B[2] = dp[2][2] + 1 = 2
    dp := make([][]int, N + 1)
    for i := 0; i < N + 1; i++ {
        dp[i] = make([]int, M + 1)
    }

    ans := 0

    for i := 1; i <= N; i++ {
        for j := 1; j <= M; j++ {
            if i == 0 || j == 0 {
                dp[i][j] = 0
            } else if A[i-1] == B[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
                ans = max(ans, dp[i][j])
            } else {
                dp[i][j] = 0
            }
        }
    }

    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
