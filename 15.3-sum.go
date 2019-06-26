/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (23.08%)
 * Likes:    3638
 * Dislikes: 399
 * Total Accepted:    534K
 * Total Submissions: 2.2M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 * 
 * Note:
 * 
 * The solution set must not contain duplicate triplets.
 * 
 * Example:
 * 
 * 
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 * 
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 * 
 * 
 */
func threeSum(nums []int) [][]int {
    N := len(nums)
    ans := [][]int{}

    if N < 3 {
        return ans
    }

    sort.Ints(nums)

    for i := 0; i < N - 2; i++ {
        if nums[i] > 0 {
            break
        }

        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        target := 0 - nums[i]

        left := i + 1
        right := N - 1

        for left < right {
            if nums[left] + nums[right] == target {
                ans = append(ans, []int{nums[i], nums[left], nums[right]})

                for left < right && nums[left] == nums[left+1] {
                    left++
                }

                for left < right && nums[right] == nums[right-1] {
                    right--
                }

                left++
                right--
            } else if nums[left] + nums[right] < target {
                left++
            } else {
                right--
            }
        }
    }

    return ans
}
