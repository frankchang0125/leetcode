/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (30.10%)
 * Total Accepted:    259.3K
 * Total Submissions: 847.4K
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given a 2D board and a word, find if the word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cell,
 * where "adjacent" cells are those horizontally or vertically neighboring. The
 * same letter cell may not be used more than once.
 *
 * Example:
 *
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * Given word = "ABCCED", return true.
 * Given word = "SEE", return true.
 * Given word = "ABCB", return false.
 *
 *
 */
// Reference: https://goo.gl/XCU2At
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	N := len(board)
	if N == 0 {
		return false
	}

	M := len(board[0])

	// Iterate each letter in 2D board.
	for y := 0; y < M; y++ {
		for x := 0; x < N; x++ {
			found := dfs(board, word, M, N, 0, x, y)
			if found {
				return true
			}
		}
	}

	return false
}

// Use DFS to visit all neighbor letters and its
// neighbor letters' neighbor letters of current letter.
//
// i: Index of matching letter of 'word'.
func dfs(board [][]byte, word string, M int, N int,
	i int, x int, y int) bool {
	if x < 0 || x >= N || y < 0 || y >= M {
		return false
	}

	if board[x][y] == word[i] {
		if i == len(word)-1 {
			// Word found, return true.
			return true
		}

		// Mark current letter as used.
		board[x][y] = '0'
		found := dfs(board, word, M, N, i+1, x-1, y) || // Go left
			dfs(board, word, M, N, i+1, x, y-1) || // Go up
			dfs(board, word, M, N, i+1, x+1, y) || // Go right
			dfs(board, word, M, N, i+1, x, y+1) // Go down

		// Reset current letter to prevent affecting other DFS.
		board[x][y] = word[i]
		return found
	}

	// Current letter is not equals to word[i].
	return false
}
