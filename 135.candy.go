/*
 * @lc app=leetcode id=135 lang=golang
 *
 * [135] Candy
 *
 * https://leetcode.com/problems/candy/description/
 *
 * algorithms
 * Hard (28.27%)
 * Likes:    502
 * Dislikes: 116
 * Total Accepted:    103.4K
 * Total Submissions: 363.1K
 * Testcase Example:  '[1,0,2]'
 *
 * There are N children standing in a line. Each child is assigned a rating
 * value.
 * 
 * You are giving candies to these children subjected to the following
 * requirements:
 * 
 * 
 * Each child must have at least one candy.
 * Children with a higher rating get more candies than their neighbors.
 * 
 * 
 * What is the minimum candies you must give?
 * 
 * Example 1:
 * 
 * 
 * Input: [1,0,2]
 * Output: 5
 * Explanation: You can allocate to the first, second and third child with 2,
 * 1, 2 candies respectively.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [1,2,2]
 * Output: 4
 * Explanation: You can allocate to the first, second and third child with 1,
 * 2, 1 candies respectively.
 * ⁠            The third child gets 1 candy because it satisfies the above two
 * conditions.
 * 
 * 
 */
func candy(ratings []int) int {
    N := len(ratings)
    dp := make([]int, N)
    dp[0] = 1

    // Scan from left to right.
    for i := 1; i < N; i++ {
        if ratings[i] > ratings[i-1] {
            dp[i] = dp[i-1] + 1
        } else {
            dp[i] = 1
        }
    }

    // Then, scan from right to left.
    for i := N - 2; i >= 0; i-- {
        if ratings[i] > ratings[i+1] {
            dp[i] = max(dp[i], dp[i+1] + 1)
        }
    }

    sum := 0
    for _, num := range dp {
        sum += num
    }
    return sum
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
