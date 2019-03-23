/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 *
 * https://leetcode.com/problems/flood-fill/description/
 *
 * algorithms
 * Easy (49.50%)
 * Total Accepted:    39.9K
 * Total Submissions: 79.6K
 * Testcase Example:  '[[1,1,1],[1,1,0],[1,0,1]]\n1\n1\n2'
 *
 *
 * An image is represented by a 2-D array of integers, each integer
 * representing the pixel value of the image (from 0 to 65535).
 *
 * Given a coordinate (sr, sc) representing the starting pixel (row and column)
 * of the flood fill, and a pixel value newColor, "flood fill" the image.
 *
 * To perform a "flood fill", consider the starting pixel, plus any pixels
 * connected 4-directionally to the starting pixel of the same color as the
 * starting pixel, plus any pixels connected 4-directionally to those pixels
 * (also with the same color as the starting pixel), and so on.  Replace the
 * color of all of the aforementioned pixels with the newColor.
 *
 * At the end, return the modified image.
 *
 * Example 1:
 *
 * Input:
 * image = [[1,1,1],[1,1,0],[1,0,1]]
 * sr = 1, sc = 1, newColor = 2
 * Output: [[2,2,2],[2,2,0],[2,0,1]]
 * Explanation:
 * From the center of the image (with position (sr, sc) = (1, 1)), all pixels
 * connected
 * by a path of the same color as the starting pixel are colored with the new
 * color.
 * Note the bottom corner is not colored 2, because it is not 4-directionally
 * connected
 * to the starting pixel.
 *
 *
 *
 * Note:
 * The length of image and image[0] will be in the range [1, 50].
 * The given starting pixel will satisfy 0  and 0 .
 * The value of each color in image[i][j] and newColor will be an integer in
 * [0, 65535].
 *
 */
// Reference: https://goo.gl/gJHHAp
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		// The starting pixel is already filled, do nothing.
		return image
	}

	dfs(image, sr, sc, image[sr][sc], newColor)
	return image
}

// Use DFS to fill all the surrounding pixels recursively.
func dfs(image [][]int, sr int, sc int, initColor int, newColor int) {
	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) {
		// Out of bound, do nothing.
		return
	}

	if image[sr][sc] != initColor {
		// Do nothing with different-color pixel.
		return
	}

	// Fill the pixel with new color.
	image[sr][sc] = newColor

	// Fill the surrounding pixels.
	dfs(image, sr-1, sc, initColor, newColor) // Go left
	dfs(image, sr, sc-1, initColor, newColor) // Go up
	dfs(image, sr+1, sc, initColor, newColor) // Go right
	dfs(image, sr, sc+1, initColor, newColor) // Go down
}
