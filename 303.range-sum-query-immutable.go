/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 *
 * https://leetcode.com/problems/range-sum-query-immutable/description/
 *
 * algorithms
 * Easy (36.20%)
 * Total Accepted:    130.4K
 * Total Submissions: 352.5K
 * Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
 *
 * Given an integer array nums, find the sum of the elements between indices i
 * and j (i ≤ j), inclusive.
 *
 * Example:
 *
 * Given nums = [-2, 0, 3, -5, 2, -1]
 *
 * sumRange(0, 2) -> 1
 * sumRange(2, 5) -> -1
 * sumRange(0, 5) -> -3
 *
 *
 *
 * Note:
 *
 * You may assume that the array does not change.
 * There are many calls to sumRange function.
 *
 *
 */
type NumArray struct {
	// sums[i]: The sum value of nums[0->i].
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))

	// Initialization, calculate every sum values for range: [0->i].
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sums[i] = nums[i]
		} else {
			sums[i] = sums[i-1] + nums[i]
		}
	}

	return NumArray{
		sums: sums,
	}
}

// This is a 'Prefix Sum' problem.
// Reference: https://goo.gl/Ez4rLb
func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sums[j]
	} else {
		// Formula:
		// 	sums[i->j] = sums[j] - sums[i-1], where j > 0.
		// 		ex: [-2, 0, 3, -5, 2, -1]
		// 		sums[2, 4] = sums[4] - sums[1]
		//				   = -2 - (-2)
		//				   = 0
		return this.sums[j] - this.sums[i-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
