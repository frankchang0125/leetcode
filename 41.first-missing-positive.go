/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 *
 * https://leetcode.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (28.01%)
 * Likes:    1575
 * Dislikes: 550
 * Total Accepted:    207.3K
 * Total Submissions: 722K
 * Testcase Example:  '[1,2,0]'
 *
 * Given an unsorted integer array, find the smallest missing positive
 * integer.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,0]
 * Output: 3
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,4,-1,1]
 * Output: 2
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: [7,8,9,11,12]
 * Output: 1
 * 
 * 
 * Note:
 * 
 * Your algorithm should run in O(n) time and uses constant extra space.
 * 
 */
// Time Complexity: O(n)
// Space Complexity: O(1)
//
// References:
//  http://bit.ly/2ZGrv4V
//  http://bit.ly/2ZCIKE7
func firstMissingPositive(nums []int) int {
    N := len(nums)

    // Total N numbers, the first missing number must in range: 1->N.
    // Therefore, numbers which < 1 or > N can be discarded.
    //
    // Try to put the number: 'num' to its correct index: 'num-1'.
    //  ex: For number: 1, put to its index: 0
    //      For number: 2, put to its index: 1
    //      For number: 3, put to its index: 2
    //      For number: 4, put to its index: 3
    //                     ...
    //      For number: N, put to its index: N-1
    //
    // Keep swapping numbers to its correct index.
    // After swapping, the first number which is not at its correct index
    // is the first missing positive number.
    //  ex: [3, 4, -1, 1] (i = 0, swap: 3 and 1)
    //   => [-1, 4, 3, 1] (i = 1, swap: 4 and 1)
    //   => [-1, 1, 3, 4] (i = 2, do nothing)
    //   => [-1, 1, 3, 3] (i = 3, do nothing)
    for i := range nums {
        for nums[i] >= 1 && nums[i] <= N && nums[i] != nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }

    // Find the first missing positive number.
    for i := 0; i < N; i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }

    return N+1
}
