/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 *
 * https://leetcode.com/problems/single-number/description/
 *
 * algorithms
 * Easy (58.59%)
 * Total Accepted:    444.2K
 * Total Submissions: 744.5K
 * Testcase Example:  '[2,2,1]'
 *
 * Given a non-empty array of integers, every element appears twice except for
 * one. Find that single one.
 *
 * Note:
 *
 * Your algorithm should have a linear runtime complexity. Could you implement
 * it without using extra memory?
 *
 * Example 1:
 *
 *
 * Input: [2,2,1]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: [4,1,2,1,2]
 * Output: 4
 *
 *
 */
func singleNumber(nums []int) int {
	result := 0

    // Use XOR to find out the single number:
    //  XOR:
    //      True  XOR True  = False
    //      True  XOR False = True
    //      False XOR True  = True
    //      False XOR False = False
    //
    // For example: [4,1,2,1,2]:
    //  0 XOR 4 = 4 (0000 XOR 0100 = 0100)
    //  4 XOR 1 = 5 (0100 XOR 0001 = 0101)
    //  5 XOR 2 = 7 (0101 XOR 0010 = 0111)
    //  7 XOR 1 = 6 (0111 XOR 0001 = 0110)
    //  6 XOR 2 = 4 (0110 XOR 0010 = 0100)
	for _, num := range nums {
		result ^= num
	}

	return result
}
