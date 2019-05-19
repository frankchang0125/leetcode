/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 *
 * https://leetcode.com/problems/number-of-1-bits/description/
 *
 * algorithms
 * Easy (43.04%)
 * Likes:    419
 * Dislikes: 372
 * Total Accepted:    256K
 * Total Submissions: 594.6K
 * Testcase Example:  '00000000000000000000000000001011'
 *
 * Write a function that takes an unsigned integer and return the number of '1'
 * bits it has (also known as the Hamming weight).
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: 00000000000000000000000000001011
 * Output: 3
 * Explanation: The input binary string 00000000000000000000000000001011 has a
 * total of three '1' bits.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 00000000000000000000000010000000
 * Output: 1
 * Explanation: The input binary string 00000000000000000000000010000000 has a
 * total of one '1' bit.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 11111111111111111111111111111101
 * Output: 31
 * Explanation: The input binary string 11111111111111111111111111111101 has a
 * total of thirty one '1' bits.
 * 
 * 
 * 
 * Note:
 * 
 * 
 * Note that in some languages such as Java, there is no unsigned integer type.
 * In this case, the input will be given as signed integer type and should not
 * affect your implementation, as the internal binary representation of the
 * integer is the same whether it is signed or unsigned.
 * In Java, the compiler represents the signed integers using 2's complement
 * notation. Therefore, in Example 3 above the input represents the signed
 * integer -3.
 * 
 * 
 * 
 * 
 * Follow up:
 * 
 * If this function is called many times, how would you optimize it?
 * 
 */
// (n - 1) & n will turn the right-most '1' bit to 0 of n. 
//  ex: 6 (0110) => (0110 - 1) & 0110 = 0101 & 0110 = 0100
//  ex: 7 (0111) => (0111 - 1) & 0111 = 0110 & 0111 = 0110
//  ex: 8 (1000) => (1000 - 1) & 1000 = 0111 & 1000 = 0000
// Thus, we can keep doing the same operations until all '1' bits
// of n have turned to 0 to count the number of '1' bits of n.
func hammingWeight(num uint32) int {
    count := 0

    for num != 0 {
        count++
        num = (num-1)&num
    }

    return count
}
