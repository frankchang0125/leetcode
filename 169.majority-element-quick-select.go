/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (52.50%)
 * Likes:    1614
 * Dislikes: 141
 * Total Accepted:    379.2K
 * Total Submissions: 722K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 * 
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,2,3]
 * Output: 3
 * 
 * Example 2:
 * 
 * 
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 * 
 * 
 */
// Since majority element is the element that appears
// more than ⌊ n/2 ⌋ times, then after we sort the array,
// the median of the sorted array must be the
// majority element.
// We can use quick select algorithm to find the
// median ([n/2]th smallest number) of the array.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func majorityElement(nums []int) int {
    // Median = kth smallest number, where k = len(nums) / 2 + 1
    // Note:
    //	For even length array, since majority element
    //  appears more than ⌊ n/2 ⌋ times:
    //		[len(nums)/2]th & [len(nums)/2+1]th smallest numbers must be same.
    //      ex: nums = [1, 2, 2, 2, 2, 3], len(nums) = 6
    //			3th smallest and 4th smallest number are same (2).
    return QuickSelect(nums, 0, len(nums) - 1, len(nums)/2+1)
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
	// numbers smaller than pivot and numbers larger than pivot
	// i is initialized to: (front - 1) since nothing has been
	// compared yet.
	i := front - 1

	// Compare all numbers (except pivot itself)
	// and swap numbers: arr[i+1] and arr[j]
	// so that all numbers which are smaller than pivot
	// will be at the left hand side of all numbers which are
	// larger than pivot.
	for j := front; j < end; j++ {
		if arr[j] < pivot {
            i++
			swap(arr, i, j)
		}
	}

	// All numbers except pivot have been compared.
	// Now all numbers: arr[front->i] are smaller than pivot
	// and all numbers: arr[i+1->end-1] are larger than pivot.
	// Swap pivot with the first number which is larger than pivot (i++)
	// so that pivot will be at the position which:
	// all numbers at the left hand side of pivot are smaller than pivot
	// and all numbers at the right hand side of pivot are larger than pivot.
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
