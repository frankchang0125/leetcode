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
// Use DFS + Greedy + Pruning to find the minimum changes
// needed to change amount.
// Reference: https://goo.gl/r66oR3
func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}

	// Sort the coins from large to small for greedy approach.
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	minChanges := amount + 1
	dfs(coins, 0, amount, 0, &minChanges)

	if minChanges == amount+1 {
		return -1
	}
	return minChanges
}

// Use DFS to find out all the possible coin changes.
// Use one type of coin for each level of DFS tree.
// s: Index of current coin to be used.
// amount: The amount of money to be changed.
// changes: Current changes so far.
// minChanges: Minimum changes so far.
func dfs(coins []int, s int, amount int, changes int, minChanges *int) {
	coin := coins[s]

	if s == len(coins)-1 {
		// If we are using last type of coins,
		// check if it can be used to change all amount of money.
		// If so, update minChanges if:
		// 	(changes so far + last type of coins required changes) is smaller than minChanges.
		if amount%coin == 0 {
			curChanges := amount / coin
			if curChanges+changes < *minChanges {
				*minChanges = curChanges + changes
			}
		}
	} else {
		maxChanges := amount / coin

		// DFS to find out all available changes (including not changing)
		// can be used for current type of coin.
		for j := maxChanges; j >= 0; j-- {
			// Pruning, if changes so far + available change is already larger than
			// minChanges, there's no need to perform further DFS.
			if changes+j >= *minChanges {
				return
			}

			// Update amount, changes and set s = s+1 to use next type of coins
			// in next-level DFS.
			dfs(coins, s+1, amount-coin*j, changes+j, minChanges)
		}
	}
}
