/*
 * @lc app=leetcode id=518 lang=golang
 *
 * [518] Coin Change 2
 *
 * https://leetcode.com/problems/coin-change-2/description/
 *
 * algorithms
 * Medium (40.86%)
 * Total Accepted:    37.3K
 * Total Submissions: 89.7K
 * Testcase Example:  '5\n[1,2,5]'
 *
 * You are given coins of different denominations and a total amount of money.
 * Write a function to compute the number of combinations that make up that
 * amount. You may assume that you have infinite number of each kind of
 * coin.
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: amount = 5, coins = [1, 2, 5]
 * Output: 4
 * Explanation: there are four ways to make up the amount:
 * 5=5
 * 5=2+2+1
 * 5=2+1+1+1
 * 5=1+1+1+1+1
 *
 *
 * Example 2:
 *
 *
 * Input: amount = 3, coins = [2]
 * Output: 0
 * Explanation: the amount of 3 cannot be made up just with coins of 2.
 *
 *
 * Example 3:
 *
 *
 * Input: amount = 10, coins = [10]
 * Output: 1
 *
 *
 *
 *
 * Note:
 *
 * You can assume that
 *
 *
 * 0 <= amount <= 5000
 * 1 <= coin <= 5000
 * the number of coins is less than 500
 * the answer is guaranteed to fit into signed 32-bit integer
 *
 *
 */
// dp[i]: Number of combinations to use all coins to make up amount 'i'.
//	ex: amount: 3, coins: [1, 2, 5]
//		iterations:
//			dp[0] = 1
//			dp[0+1] += dp[0] => dp[1] += 1
//			dp[1+1] += dp[1] => dp[2] += 1
//			dp[2+1] += dp[2] => dp[3] += 1
//			dp[0+2] += dp[0] => dp[2] += 1
//			dp[1+2] += dp[1] => dp[3] += 1
//
//		=> dp[0] = 1
//
//		=> dp[1] = sum of:
//			=> dp[0+1] = dp[0]
//					   = 1
//		   = 1
//
//		=> dp[2] = sum of:
//			=> dp[1+1] = dp[1]
//					   = 1
//			=> dp[0+2] = dp[0]
//					   = 1
//		   = 2
//
//		=> dp[3] = sum of:
//			=> dp[2+1] = dp[2]
//					   = 1 (Was still: '1' at j = 2, i = 0 (i.e. coins[i] = 1) iteration.)
//			=> dp[1+2] = dp[1]
//					   = 1
//		   = 2
//
// References:
//	http://bit.ly/2UDXm88
//	http://bit.ly/2UCAF4n
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount-coins[i]; j++ {
			dp[j+coins[i]] += dp[j]
		}
	}

	return dp[amount]
}
