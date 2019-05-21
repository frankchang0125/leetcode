/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 *
 * https://leetcode.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (47.59%)
 * Likes:    1974
 * Dislikes: 165
 * Total Accepted:    369.1K
 * Total Submissions: 774.8K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * Find the kth largest element in an unsorted array. Note that it is the kth
 * largest element in the sorted order, not the kth distinct element.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,2,1,5,6,4] and k = 2
 * Output: 5
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,2,3,1,2,4,5,5,6] and k = 4
 * Output: 4
 * 
 * Note: 
 * You may assume k is always valid, 1 ≤ k ≤ array's length.
 * 
 */
func findKthLargest(nums []int, k int) int {
    return QuickSelect(nums, 0, len(nums) - 1, k)
}

// Quick select steps:
//	1. Find a pivot which separates numbers into two parts:
//		a. Left: Numbers smaller than pivot.
//		b. Right: Numbers larger than pivot.
//	2. Keep repeating Step 1. until we can find a pivot
//	   which is the kth largest number.
//	   (i.e. pivot's index == k - 1)
func QuickSelect(arr []int, front int, end int, k int) int {
	if len(arr) == 0 || len(arr) < k {
		return -1
	}

	// Only one number in this interval, just return it.
	if front == end {
		return arr[front]
	}

	// Find the index of pivot.
	pivotIdx := partition(arr, front, end)

	// Keep repeating until we find a pivot which is the kth smallest number.
	for pivotIdx != k - 1 {
		if pivotIdx > k - 1 {
			// Pivot is at the right hand side of kth smallest number,
			// search left part.
			return QuickSelect(arr, front, pivotIdx - 1, k);
		} else {
			// Pivot is at the left hand side of kth smallest number,
			// search right part.
			return QuickSelect(arr, pivotIdx + 1, end, k)
		}
	}

	return arr[pivotIdx]
}

func partition(arr []int, front int, end int) int {
	// Pivot can be any number in array, we just choose the last number here.
	pivot := arr[end]

	// i will be the final index of pivot,
	// which separates numbers into two parts:
	// numbers larger than pivot and numbers smaller than pivot
	// i is initialized to: (front - 1) since nothing has been
	// compared yet.
	i := front - 1

	// Compare all numbers (except pivot itself)
	// and swap numbers: arr[i+1] and arr[j]
	// so that all numbers which are larger than pivot
	// will be at the left hand side of all numbers which are
	// smaller than pivot.
	for j := front; j < end; j++ {
		if arr[j] > pivot {
			i++
			swap(arr, i, j)
		}
	}

	// All numbers except pivot have been compared.
	// Now all numbers: arr[front->i] are larger than pivot
	// and all numbers: arr[i+1->end-1] are smaller than pivot.
	// Swap pivot with the first number which is smaller than pivot (i++)
	// so that pivot will be at the position which:
	// all numbers at the left hand side of pivot are larger than pivot
	// and all numbers at the right hand side of pivot are smaller than pivot.
	i++
	swap(arr, i, end)

	// Return pivot's index.
	return i
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
