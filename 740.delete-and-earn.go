/*
 * @lc app=leetcode id=740 lang=golang
 *
 * [740] Delete and Earn
 *
 * https://leetcode.com/problems/delete-and-earn/description/
 *
 * algorithms
 * Medium (44.78%)
 * Total Accepted:    19.8K
 * Total Submissions: 43.6K
 * Testcase Example:  '[3,4,2]'
 *
 * Given an array nums of integers, you can perform operations on the array.
 *
 * In each operation, you pick any nums[i] and delete it to earn nums[i]
 * points. After, you must delete every element equal to nums[i] - 1 or nums[i]
 * + 1.
 *
 * You start with 0 points. Return the maximum number of points you can earn by
 * applying such operations.
 *
 * Example 1:
 *
 *
 * Input: nums = [3, 4, 2]
 * Output: 6
 * Explanation:
 * Delete 4 to earn 4 points, consequently 3 is also deleted.
 * Then, delete 2 to earn 2 points. 6 total points are earned.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2, 2, 3, 3, 3, 4]
 * Output: 9
 * Explanation:
 * Delete 3 to earn 3 points, deleting both 2's and the 4.
 * Then, delete 3 again to earn 3 points, and 3 again to earn 3 points.
 * 9 total points are earned.
 *
 *
 *
 *
 * Note:
 *
 *
 * The length of nums is at most 20000.
 * Each element nums[i] is an integer in the range [1, 10000].
 *
 *
 *
 *
 */
// Reference: http://bit.ly/2D16sky
func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
    }
    
    maxVal := math.MinInt32
    for i := 0; i < len(nums); i++ {
        if nums[i] > maxVal {
            maxVal = nums[i]
        }
    }

    newNums := make([]int, maxVal+1)
    newNums[0]= 0

    for i := 0; i < len(nums); i++ {
        newNums[nums[i]] += nums[i]
    }
    fmt.Println(newNums)

	return rob(newNums)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp1 := nums[0]
	dp2 := 0

	for i := 1; i < len(nums); i++ {
		if i == 1 {
			dp2 = max(nums[i], dp1)
		} else {
			dp2 = max(dp2+nums[i], dp1)
		}
		dp1, dp2 = dp2, dp1
	}

	return dp1
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
