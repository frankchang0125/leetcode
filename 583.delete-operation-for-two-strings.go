/*
 * @lc app=leetcode id=583 lang=golang
 *
 * [583] Delete Operation for Two Strings
 *
 * https://leetcode.com/problems/delete-operation-for-two-strings/description/
 *
 * algorithms
 * Medium (44.76%)
 * Likes:    694
 * Dislikes: 19
 * Total Accepted:    31.5K
 * Total Submissions: 69.9K
 * Testcase Example:  '"sea"\n"eat"'
 *
 * 
 * Given two words word1 and word2, find the minimum number of steps required
 * to make word1 and word2 the same, where in each step you can delete one
 * character in either string.
 * 
 * 
 * Example 1:
 * 
 * Input: "sea", "eat"
 * Output: 2
 * Explanation: You need one step to make "sea" to "ea" and another step to
 * make "eat" to "ea".
 * 
 * 
 * 
 * Note:
 * 
 * The length of given words won't exceed 500.
 * Characters in given words can only be lower-case letters.
 * 
 * 
 */
// This is the variation of "Longest Common Subsequence" problem.
// Reference: http://bit.ly/2FzXDPV
func minDistance(word1 string, word2 string) int {
    N := len(word1)
    M := len(word2)

    // Let A = word1 and B = word2.
    // dp[i][j] = Length of longest common subsequence of A[0->i-1] and B[0->i-1].
    // dp[i][j]:
    //  1. If i == 0 or j == 0:
    //      => Empty substring, common subsequence count = 0.
    //  2. If A[i-1] == B[i-1]:
    //      => dp[i-1][j-1] + 1
    //  3. If A[i-1] != B[i-1]:
    //      => Since we are looking for subsequence (not substring),
    //         choose the longest common subsequence between (A[0->i-2], B[0->j-1]) and (A[0->i-1], B[0->j-2]).
    //      => max(dp[i-1][j], dp[j-1][i])
    //  ex: A = "abedf", B = "aaeda":
    //      => dp[0][0] = 0
    //      => dp[0][1] = 0
    //      => dp[0][2] = 0
    //      => dp[0][3] = 0
    //      => dp[0][4] = 0
    //      => dp[0][5] = 0
    //      => dp[1][0] = 0
    //      => dp[1][1]: A[0] == B[0] = dp[0][0] + 1 = 1
    //      => dp[1][2]: A[0] == B[1] = dp[0][1] + 1 = 1
    //      => dp[1][3]: A[0] != B[2] = max("" + "aae", "a" + "aa") = max(dp[0][3], dp[1][2]) = 1
    //      => dp[1][4]: A[0] != B[3] = max("" + "aaed", "a" + "aae") = max(dp[0][4], dp[1][3]) = 1
    //      => dp[1][5]: A[0] == B[4] = dp[0][4] + 1 = 1
    //      => dp[2][0] = 0
    //      => dp[2][1]: A[1] != B[0] = max("a" + "a", "ab" + "") = max(dp[1][1], dp[2][0]) = 1
    //      => dp[2][2]: A[1] != B[1] = max("a" + "aa", "ab" + "a") = max(dp[1][2], dp[2][1]) = 1
    //      => dp[2][3]: A[1] != B[2] = max("a" + "aae", "ab" + "aa") = max(dp[1][3], dp[2][2]) = 1
    //      => dp[2][4]: A[1] != B[3] = max("a" + "aaed", "ab" + "aae") = max(dp[1][4], dp[2][3]) = 1
    //      => dp[2][5]: A[1] != B[4] = max("a" + "aaeda", "ab" + "aaed") = max(dp[1][5], dp[2][4]) = 1
    //      => dp[3][0] = 0
    //      => dp[3][1]: A[2] != B[0] = max("ab" + "a", "abe" + "") = max(dp[2][1], dp[3][0]) = 1
    //      => dp[3][2]: A[2] == B[1] = max("ab" + "aa", "abe" + "a") = max(dp[2][2], dp[3][1]) = 1
    //      => dp[3][3]: A[2] == B[2] = dp[2][2] + 1 = 2
    //      => dp[3][4]: A[2] != B[3] = max("ab" + "aaed", "abe" + "aae") = max(dp[2][4], dp[3][3]) = 2
    //      => dp[3][5]: A[2] != B[4] = max("ab" + "aaeda", "abe" + "aaed") = max(dp[2][5], dp[3][4]) = 2
    //      => dp[4][0] = 0
    //      => dp[4][1]: A[3] != B[0] = max("abe" + "a", "abed" + "") = max(dp[3][1], dp[4][0]) = 1
    //      => dp[4][2]: A[3] != B[1] = max("abe" + "aa", "abed" + "a") = max(dp[3][2], dp[4][1]) = 1
    //      => dp[4][3]: A[3] == B[2] = max("abe" + "aae", "abed" + "aa") = max(dp[3][3], dp[4][2]) = 2
    //      => dp[4][4]: A[3] == B[3] = dp[3][3] + 1 = 3
    //      => dp[4][5]: A[3] != B[4] = max("abe" + "aaeda", "abed" + "aaed") = max(dp[3][5], dp[4][4]) = 3
    //      => dp[5][0] = 0
    //      => dp[5][1]: A[4] != B[0] = max("abed" + "a", "abedf" + "") = max(dp[4][1], dp[5][0]) = 1
    //      => dp[5][2]: A[4] != B[1] = max("abed" + "aa", "abedf" + "a") = max(dp[4][2], dp[5][1]) = 1
    //      => dp[5][3]: A[4] != B[2] = max("abed" + "aae", "abedf" + "aa") = max(dp[4][3], dp[5][2]) = 2
    //      => dp[5][4]: A[4] != B[3] = max("abed" + "aaed", "abedf" + "aae") = max(dp[4][4], dp[5][3]) = 3
    //      => dp[5][5]: A[4] != B[4] = max("abed" + "aaeda", "abedf" + "aaed") = max(dp[4][5], dp[5][4]) = 3
    dp := make([][]int, N + 1)
    for i := 0; i < N + 1; i++ {
        dp[i] = make([]int, M + 1)
    }

    for i := 1; i <= N; i++ {
        for j := 1; j <= M; j++ {
            if i == 0 || j == 0 {
                dp[i][j] = 0
            } else if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }

    longestCommonSubsequenceLength := dp[N][M]

    // Since we have found the longest common subsequence's length,
    // remove the rest of characters from both word1 and word2
    // then word1 and word2 will be same.
    ans := len(word1) - longestCommonSubsequenceLength
    ans += len(word2) - longestCommonSubsequenceLength
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
