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
// Space Complexity: O(n)
//
// References:
//  http://bit.ly/2ZGrv4V
//  http://bit.ly/2ZCIKE7
func firstMissingPositive(nums []int) int {
    N := len(nums)
    m := map[int]struct{}{}

    // Total N numbers, the first missing number must in range: 1->N.
    // Therefore, numbers which < 1 or > N can be discarded.
    for _, num := range nums {
        if num < 1 || num > N {
            continue
        }
        m[num] = struct{}{}
    }

    // Find the first missing positive number.
    for i := 1; i <= N; i++ {
        if _, ok := m[i]; !ok {
            return i
        }
    }

    return N+1
}
