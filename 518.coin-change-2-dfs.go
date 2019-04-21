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
// Note:
// 	Use 39. combination sum (DFS) solution,
// 	however, it will exceed time limitation!!
//
// References:
//	http://bit.ly/2UDXm88
//	http://bit.ly/2UCAF4n
func change(amount int, coins []int) int {
	if len(coins) == 0 {
        if amount == 0 {
            return 1
        }
		return 0
    }
    
	sort.Sort(sort.IntSlice(coins))
	ans := 0
	dfs(coins, amount, 0, 0, &ans)
	return ans
}

// s: Start index of coins to be added to sum.
func dfs(coins []int, amount int, s int, sum int, ans *int) {
	if sum == amount {
		*ans++
		return
	}

	for i := s; i < len(coins); i++ {
		if sum+coins[i] > amount {
			// Stop further iterations if sum+coins[i]
			// has already exceeded the value of target.
			// (We've sorted coins already.)
			return
		}

		// Different from pure combination problem,
		// we can use current coins repeatedly,
		// thus pass 'i' instead of 'i+1' to next DFS.
		dfs(coins, amount, i, sum+coins[i], ans)
	}
}
