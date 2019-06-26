/*
 * @lc app=leetcode id=801 lang=golang
 *
 * [801] Minimum Swaps To Make Sequences Increasing
 *
 * https://leetcode.com/problems/minimum-swaps-to-make-sequences-increasing/description/
 *
 * algorithms
 * Medium (33.88%)
 * Total Accepted:    11.9K
 * Total Submissions: 35.1K
 * Testcase Example:  '[1,3,5,4]\n[1,2,3,7]'
 *
 * We have two integer sequences A and B of the same non-zero length.
 *
 * We are allowed to swap elements A[i] and B[i].  Note that both elements are
 * in the same index position in their respective sequences.
 *
 * At the end of some number of swaps, A and B are both strictly increasing.
 * (A sequence is strictly increasing if and only if A[0] < A[1] < A[2] < ... <
 * A[A.length - 1].)
 *
 * Given A and B, return the minimum number of swaps to make both sequences
 * strictly increasing.  It is guaranteed that the given input always makes it
 * possible.
 *
 *
 * Example:
 * Input: A = [1,3,5,4], B = [1,2,3,7]
 * Output: 1
 * Explanation:
 * Swap A[3] and B[3].  Then the sequences are:
 * A = [1, 3, 5, 7] and B = [1, 2, 3, 4]
 * which are both strictly increasing.
 *
 *
 * Note:
 *
 *
 * A, B are arrays with the same length, and that length will be in the range
 * [1, 1000].
 * A[i], B[i] are integer values in the range [0, 2000].
 *
 *
 */
func minSwap(A []int, B []int) int {
	keep := make([]int, len(A))
	swap := make([]int, len(A))

	// Initialization.
	for i := 0; i < len(A); i++ {
		keep[i] = math.MaxInt32
		swap[i] = math.MaxInt32
	}

	keep[0] = 0
	swap[0] = 1

	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] && B[i] > B[i-1] {
			keep[i] = keep[i-1]
			swap[i] = swap[i-1] + 1
		}

		if B[i] > A[i-1] && A[i] > B[i-1] {
			swap[i] = min(swap[i], keep[i-1]+1)
			keep[i] = min(keep[i], swap[i-1])
		}
	}

	return min(keep[len(keep)-1], swap[len(swap)-1])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
