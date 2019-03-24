/*
 * @lc app=leetcode id=464 lang=golang
 *
 * [464] Can I Win
 *
 * https://leetcode.com/problems/can-i-win/description/
 *
 * algorithms
 * Medium (26.68%)
 * Total Accepted:    33.5K
 * Total Submissions: 124.1K
 * Testcase Example:  '10\n11'
 *
 * In the "100 game," two players take turns adding, to a running total, any
 * integer from 1..10. The player who first causes the running total to reach
 * or exceed 100 wins.
 *
 * What if we change the game so that players cannot re-use integers?
 *
 * For example, two players might take turns drawing from a common pool of
 * numbers of 1..15 without replacement until they reach a total >= 100.
 *
 * Given an integer maxChoosableInteger and another integer desiredTotal,
 * determine if the first player to move can force a win, assuming both players
 * play optimally.
 *
 * You can always assume that maxChoosableInteger will not be larger than 20
 * and desiredTotal will not be larger than 300.
 *
 *
 * Example
 *
 * Input:
 * maxChoosableInteger = 10
 * desiredTotal = 11
 *
 * Output:
 * false
 *
 * Explanation:
 * No matter which integer the first player choose, the first player will lose.
 * The first player can choose an integer from 1 up to 10.
 * If the first player choose 1, the second player can only choose integers
 * from 2 up to 10.
 * The second player will win by choosing 10 and get a total = 11, which is >=
 * desiredTotal.
 * Same with other integers chosen by the first player, the second player will
 * always win.
 *
 *
 */
// Use recursion + memoization to find out whether we can we the game or not.
//
// In order to prevent duplicate combinations been calculated again,
// we use a hash map to record the win result of every combinations of numbers been picked.
// (i.e. memoization.)
//
// To do so, use an integer: 'state' to record which numbers are been picked for each round.
//  If state[i-1] == 1, number 'i' is been picked.
//  If state[i-1] == 0, number 'i' is not been picked yet.
//  ex: state = 5 (101): number '1' and number '3' are been picked.
//
// Therefore, for games:
//  Player 1 first picked number: '1', than Player 2 picked number: '2' and
//  Player 1 first picked number: '2', than Player 2 picked number: '1'
// Both game win results will be same, and we can use 'state' to represent both combinations.
// ('state' for both combinations will all be: '3'.)
//
// References:
//  https://goo.gl/TPEfhX
//  https://goo.gl/VcNUK3
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	// If desired total number is larger than
	// the total sum amount of: 1 -> maxChoosableInteger,
	// there's no way for both players to win.
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	// If desired total number is less than zero,
	// no way to win.
	if desiredTotal < 0 {
		return false
	}

	state := 0
	mem := map[int]bool{}
	return canWin(maxChoosableInteger, desiredTotal, state, mem)
}

func canWin(max int, total int, state int, mem map[int]bool) bool {
	// Try to get the result from memoization.
	if m, ok := mem[state]; ok {
		return m
	}

	// Try to pick a number from: 1 -> total to see if I can win.
	for i := 1; i <= max; i++ {
		if state&(1<<uint(i-1)) > 0 {
			// Number 'i' is already been picked, skip.
			continue
		}

		// Number 'i' picked, win condition:
		//  1. (total - i) <= 0, or
		//  2. Opponent cannot win.
		//     (with remaining total and number 'i' been picked.)
		if total-i <= 0 || !canWin(max, total-i, state|1<<uint32(i-1), mem) {
			// Memoization.
			mem[state] = true
			return true
		}
	}

	// Can find anyway to win, memorize the result and return false.
	mem[state] = false
	return false
}
