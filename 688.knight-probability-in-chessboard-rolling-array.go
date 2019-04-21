/*
 * @lc app=leetcode id=688 lang=golang
 *
 * [688] Knight Probability in Chessboard
 *
 * https://leetcode.com/problems/knight-probability-in-chessboard/description/
 *
 * algorithms
 * Medium (42.51%)
 * Total Accepted:    18.3K
 * Total Submissions: 42.1K
 * Testcase Example:  '3\n2\n0\n0'
 *
 * On an NxN chessboard, a knight starts at the r-th row and c-th column and
 * attempts to make exactly K moves. The rows and columns are 0 indexed, so the
 * top-left square is (0, 0), and the bottom-right square is (N-1, N-1).
 *
 * A chess knight has 8 possible moves it can make, as illustrated below. Each
 * move is two squares in a cardinal direction, then one square in an
 * orthogonal direction.
 *
 *
 *
 *
 *
 *
 *
 * Each time the knight is to move, it chooses one of eight possible moves
 * uniformly at random (even if the piece would go off the chessboard) and
 * moves there.
 *
 * The knight continues moving until it has made exactly K moves or has moved
 * off the chessboard. Return the probability that the knight remains on the
 * board after it has stopped moving.
 *
 *
 *
 * Example:
 *
 *
 * Input: 3, 2, 0, 0
 * Output: 0.0625
 * Explanation: There are two moves (to (1,2), (2,1)) that will keep the knight
 * on the board.
 * From each of those positions, there are also two moves that will keep the
 * knight on the board.
 * The total probability the knight stays on the board is 0.0625.
 *
 *
 *
 *
 * Note:
 *
 *
 * N will be between 1 and 25.
 * K will be between 0 and 100.
 * The knight always initially starts on the board.
 *
 *
 */
// Reference: http://bit.ly/2UAys8G
func knightProbability(N int, K int, r int, c int) float64 {
	// dp[y][x]: Number of moves to move knight which still remain
	//           on board 's' times from point: (y, x)
	//
	// If y < 0 || y >= N || x < 0 || x >= N, out of board,
	// number of moves = 0.
	//
	// cur[y][x] = dp[y-2][x-1] +
	//             dp[y-1][x-2] +
	//             dp[y+1][x-2] +
	//             dp[y+2][x-1] +
	//             dp[y+2][x+1] +
	//             dp[y+1][x+2] +
	//             dp[y-1][x+2] +
	//             dp[y-2][x+1]
	dp := make([][]float64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]float64, N)
	}

	// directions of (y, x)
	directions := [][]int{
		{-2, -1},
		{-1, -2},
		{1, -2},
		{2, -1},
		{2, 1},
		{1, 2},
		{-1, 2},
		{-2, 1},
	}

	for s := 0; s <= K; s++ {
		cur := make([][]float64, N)
		for i := 0; i < N; i++ {
			cur[i] = make([]float64, N)
		}

		for y := 0; y < N; y++ {
			for x := 0; x < N; x++ {
				if s == 0 {
					cur[y][x] = 1
					continue
				}

				for _, d := range directions {
					newY := y + d[0]
					newX := x + d[1]

					if newY < 0 || newY >= N || newX < 0 || newX >= N {
						cur[y][x] += 0
					} else {
						cur[y][x] += dp[newY][newX]
					}
				}
			}
		}

		cur, dp = dp, cur
	}

	total := math.Pow(8, float64(K))
	return float64(dp[r][c]) / total
}
