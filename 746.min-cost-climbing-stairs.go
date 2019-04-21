/*
 * @lc app=leetcode id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 *
 * https://leetcode.com/problems/min-cost-climbing-stairs/description/
 *
 * algorithms
 * Easy (45.55%)
 * Total Accepted:    72.1K
 * Total Submissions: 155K
 * Testcase Example:  '[0,0,0,0]'
 *
 *
 * On a staircase, the i-th step has some non-negative cost cost[i] assigned (0
 * indexed).
 *
 * Once you pay the cost, you can either climb one or two steps. You need to
 * find minimum cost to reach the top of the floor, and you can either start
 * from the step with index 0, or the step with index 1.
 *
 *
 * Example 1:
 *
 * Input: cost = [10, 15, 20]
 * Output: 15
 * Explanation: Cheapest is start on cost[1], pay that cost and go to the
 * top.
 *
 *
 *
 * Example 2:
 *
 * Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
 * Output: 6
 * Explanation: Cheapest is start on cost[0], and only step on 1s, skipping
 * cost[3].
 *
 *
 *
 * Note:
 *
 * cost will have a length in the range [2, 1000].
 * Every cost[i] will be an integer in the range [0, 999].
 *
 *
 */
// Reference: http://bit.ly/2VcC417
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}

	if len(cost) == 1 {
		return cost[0]
	}

	// dp[i]: Minimum cost to climb to i-th stair.
	// dp[i] = min(Climb from last stair,
	//			   Climb from second to the last stair)
	//		   + Cost of this stair
	//		 = min(dp[i-1], dp[i-2]) + cost[i]
	dp := make([]int, len(cost))

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	// Climb to the top of floor, choose the minimum costs between
	// climbing from last stair or from second to the last stair.
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
