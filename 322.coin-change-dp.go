/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 *
 * https://leetcode.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (28.66%)
 * Total Accepted:    168.2K
 * Total Submissions: 575K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * You are given coins of different denominations and a total amount of money
 * amount. Write a function to compute the fewest number of coins that you need
 * to make up that amount. If that amount of money cannot be made up by any
 * combination of the coins, return -1.
 *
 * Example 1:
 *
 *
 * Input: coins = [1, 2, 5], amount = 11
 * Output: 3
 * Explanation: 11 = 5 + 5 + 1
 *
 * Example 2:
 *
 *
 * Input: coins = [2], amount = 3
 * Output: -1
 *
 *
 * Note:
 * You may assume that you have an infinite number of each kind of coin.
 *
 */
// Use dynamic programming to find the minimum changes
// needed to change amount.
// References:
//  https://goo.gl/r66oR3
//  https://goo.gl/6LiwoH
//  https://goo.gl/H8nftU
func coinChange(coins []int, amount int) int {
	// dp[i]: The minimum changes needed for money: i.
	dp := make([]int, amount+1)
	dp[0] = 0

	// Initialize dynamic programming boundary.
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for m := 1; m <= amount; m++ {
		// Reset minChanges.
		minChanges := amount + 1

		// For money: m
		// dp[m] = Minimum changes needed for money: m - coins[j] + 1.
		// i.e. dp[m] = min(m - coins[j]) + 1, where j = 0 -> len(coins)-1.
		for _, c := range coins {
			change := m - c

			if change < 0 {
				// Cannot make change, continue.
				continue
			}

			// Update minChanges.
			if dp[change] < minChanges {
				minChanges = dp[change]
			}
		}

		dp[m] = minChanges + 1
	}

	// Nothing can be changed, return -1.
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
