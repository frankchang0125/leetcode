/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 *
 * https://leetcode.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (52.36%)
 * Total Accepted:    205.1K
 * Total Submissions: 382.9K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * Given two arrays, write a function to compute their intersection.
 * 
 * Example 1:
 * 
 * 
 * Input: nums1 = [1,2,2,1], nums2 = [2,2]
 * Output: [2]
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * Output: [9,4]
 * 
 * 
 * Note:
 * 
 * 
 * Each element in the result must be unique.
 * The result can be in any order.
 * 
 * 
 * 
 * 
 */
func intersection(nums1 []int, nums2 []int) []int {
    ans := []int{}
    
    m := map[int]bool{}
        
    for i := 0; i < len(nums1); i++ {
        m[nums1[i]] = true
    }
        
    for i := 0; i < len(nums2); i++ {
        if _, ok := m[nums2[i]]; ok {
            ans = append(ans, nums2[i])
            delete(m, nums2[i])
        }
    }
    
    return ans 
}
