/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 *
 * https://leetcode.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (35.62%)
 * Likes:    1087
 * Dislikes: 2764
 * Total Accepted:    366.2K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as
 * one sorted array.
 * 
 * Note:
 * 
 * 
 * The number of elements initialized in nums1 and nums2 are m and n
 * respectively.
 * You may assume that nums1 has enough space (size that is greater or equal to
 * m + n) to hold additional elements from nums2.
 * 
 * 
 * Example:
 * 
 * 
 * Input:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 * 
 * Output: [1,2,2,3,5,6]
 * 
 * 
 */
func merge(nums1 []int, m int, nums2 []int, n int)  {
    N := len(nums1)
    if m + n > N {
        return
    }

    for i := (m + n) - 1; i >= 0; i-- {
        if m > 0 && n > 0 {
            if nums1[m-1] >= nums2[n-1] {
                nums1[i] = nums1[m-1]
                m--
            } else {
                nums1[i] = nums2[n-1]
                n--
            }
        } else if n > 0 && m == 0 {
            nums1[i] = nums2[n-1]
            n--
        } else if m > 0 && n == 0 {
            nums1[i] = nums1[m-1]
            m--
        } else {
            // Should not happen.
            return
        }
    } 
}
