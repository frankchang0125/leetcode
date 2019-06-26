import "fmt"

/*
 * @lc app=leetcode id=174 lang=golang
 *
 * [174] Dungeon Game
 *
 * https://leetcode.com/problems/dungeon-game/description/
 *
 * algorithms
 * Hard (26.26%)
 * Total Accepted:    63.1K
 * Total Submissions: 236.8K
 * Testcase Example:  '[[-2,-3,3],[-5,-10,1],[10,30,-5]]'
 *
 * table.dungeon, .dungeon th, .dungeon td {
 * ⁠ border:3px solid black;
 * }
 *
 * ⁠.dungeon th, .dungeon td {
 * ⁠   text-align: center;
 * ⁠   height: 70px;
 * ⁠   width: 70px;
 * }
 *
 * The demons had captured the princess (P) and imprisoned her in the
 * bottom-right corner of a dungeon. The dungeon consists of M x N rooms laid
 * out in a 2D grid. Our valiant knight (K) was initially positioned in the
 * top-left room and must fight his way through the dungeon to rescue the
 * princess.
 *
 * The knight has an initial health point represented by a positive integer. If
 * at any point his health point drops to 0 or below, he dies immediately.
 *
 * Some of the rooms are guarded by demons, so the knight loses health
 * (negative integers) upon entering these rooms; other rooms are either empty
 * (0's) or contain magic orbs that increase the knight's health (positive
 * integers).
 *
 * In order to reach the princess as quickly as possible, the knight decides to
 * move only rightward or downward in each step.
 *
 *
 *
 * Write a function to determine the knight's minimum initial health so that he
 * is able to rescue the princess.
 *
 * For example, given the dungeon below, the initial health of the knight must
 * be at least 7 if he follows the optimal path RIGHT-> RIGHT -> DOWN ->
 * DOWN.
 *
 *
 *
 *
 * -2 (K)
 * -3
 * 3
 *
 *
 * -5
 * -10
 * 1
 *
 *
 * 10
 * 30
 * -5 (P)
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * The knight's health has no upper bound.
 * Any room can contain threats or power-ups, even the first room the knight
 * enters and the bottom-right room where the princess is imprisoned.
 *
 *
 */
func calculateMinimumHP(dungeon [][]int) int {
    M := len(dungeon)
    if M == 0 {
        return 0
    }    

    N := len(dungeon[0])
    if N == 0 {
        return 0
    }

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
	}

	for y := 0; y < M; y++ {
		for x := 0; x < N; x++ {
			if x == 0 && y == 0 {
				dp[y][x] = dungeon[0][0]
			} else if x == 0 {
				dp[y][x] = dungeon[y][0] + dp[y-1][0]
			} else if y == 0 {
				dp[y][x] = dungeon[0][x] + dp[0][x-1]
			} else {
                dp[y][x] = max(1,
                    min(dungeon[y-1][x], dungeon[y][x-1])-dungeon[y][x])
			}
		}
	}

	fmt.Println(dp)

	return dp[M-1][M-1] + 1
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
