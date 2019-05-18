/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (34.62%)
 * Likes:    805
 * Dislikes: 97
 * Total Accepted:    221.5K
 * Total Submissions: 635.5K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 * 
 * 
 * Integers in each row are sorted from left to right.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * Output: false
 * 
 */
// Search from right-upper corner to left-bottom corner.
// Let the number at right-upper corner: n.
//
// If the target number = n:
//  Found the target number, return true.
// If the target number < n:
//  Since all numbers in the same column below n are greater than n,
//  its impossible that target number to be exists in the same column of n.
//  We can skip the entire column.
// If the target number > n:
//  Since all numbers in the same row left than n are smaller than n,
//  its impossible that target number to be exists in the same row of n.
//  We can skip the entire row.
func searchMatrix(matrix [][]int, target int) bool {
    M := len(matrix)
    if M == 0 {
        return false
    }

    N := len(matrix[0])
    if N == 0 {
        return false
    }

    y := 0
    x := N - 1

    for x >= 0 && y < M {
        if matrix[y][x] == target {
            return true
        } else if target < matrix[y][x] {
            // Skip the entire column.
            x--
        } else {
            // Skip the entire row.
            y++
        }
    }

    return false
}
