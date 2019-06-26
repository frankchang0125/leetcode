/*
 * @lc app=leetcode id=787 lang=golang
 *
 * [787] Cheapest Flights Within K Stops
 *
 * https://leetcode.com/problems/cheapest-flights-within-k-stops/description/
 *
 * algorithms
 * Medium (33.48%)
 * Total Accepted:    33.1K
 * Total Submissions: 97.2K
 * Testcase Example:  '3\n[[0,1,100],[1,2,100],[0,2,500]]\n0\n2\n1'
 *
 * There are n cities connected by m flights. Each fight starts from city u and
 * arrives at v with a price w.
 *
 * Now given all the cities and flights, together with starting city src and
 * the destination dst, your task is to find the cheapest price from src to dst
 * with up to k stops. If there is no such route, output -1.
 *
 *
 * Example 1:
 * Input:
 * n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
 * src = 0, dst = 2, k = 1
 * Output: 200
 * Explanation:
 * The graph looks like this:
 *
 *
 * The cheapest price from city 0 to city 2 with at most 1 stop costs 200, as
 * marked red in the picture.
 *
 *
 * Example 2:
 * Input:
 * n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
 * src = 0, dst = 2, k = 0
 * Output: 500
 * Explanation:
 * The graph looks like this:
 *
 *
 * The cheapest price from city 0 to city 2 with at most 0 stop costs 500, as
 * marked blue in the picture.
 *
 * Note:
 *
 *
 * The number of nodes n will be in range [1, 100], with nodes labeled from 0
 * to n - 1.
 * The size of flights will be in range [0, n * (n - 1) / 2].
 * The format of each flight will be (src, dst, price).
 * The price of each flight will be in the range [1, 10000].
 * k is in the range of [0, n - 1].
 * There will not be any duplicated flights or self cycles.
 *
 *
 */
// Reference: http://bit.ly/2UBlN5s
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	// dp[k][v]: Cheapest price from source city 'src' to city 'v' with at most 'k' stops.
	//
	// For 0 stops:
	//  dp[0][v] = cost[src][v] = flight price, if source city 'src' and city 'v' is connected.
	//  dp[0][v] = cost[src][v] = unlimited (math.MaxInt32), if source city 'src' and city 'v' are not connected.
	//
	// dp[k][v] = min(dp[k-1][u] + cost[u][v], dp[k-1][src][v])
	//  i.e. Minimum price of:
	//          1. From source city 'src' to city 'u' with at most 'k-1' stops + cost from city 'u' to city 'v'.
	//          2. From source city 'src' to city 'v' with at most 'k-1' stops.
	// For cost:
	//  If city 'u' and city 'v' are same, cost[u][v] = unlimited (math.MaxInt32).
	//  If city 'u' and city 'v' are connected, cost[u][v] = flight price.
	//  If city 'u' and city 'v' are not connected, cost[u][v] = unlimited (math.MaxInt32).
	dp := make([][]int, K+1)
	for i := 0; i < K+1; i++ {
		dp[i] = make([]int, n)
	}

	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
	}

	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			cost[y][x] = math.MaxInt32
		}
	}

	for _, flight := range flights {
		u, v, price := flight[0], flight[1], flight[2]
		cost[u][v] = price
	}

	for k := 0; k <= K; k++ {
		for v := 0; v < n; v++ {
			if k == 0 {
				dp[k][v] = cost[src][v]
				continue
			}

			for u := 0; u < n; u++ {
				if u == v {
                    // Cannot fly to the same city.
					continue
				}
				dp[k][v] = min(dp[k-1][u]+cost[u][v], dp[k-1][v])
			}
		}
    }
    
	if dp[K][dst] == math.MaxInt32 {
		return -1
	}
	return dp[K][dst]
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
