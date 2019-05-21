/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 *
 * https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
 *
 * algorithms
 * Easy (46.18%)
 * Likes:    292
 * Dislikes: 129
 * Total Accepted:    31.8K
 * Total Submissions: 68.8K
 * Testcase Example:  '["KthLargest","add","add","add","add","add"]\n[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]'
 *
 * Design a class to find the kth largest element in a stream. Note that it is
 * the kth largest element in the sorted order, not the kth distinct element.
 * 
 * Your KthLargest class will have a constructor which accepts an integer k and
 * an integer array nums, which contains initial elements from the stream. For
 * each call to the method KthLargest.add, return the element representing the
 * kth largest element in the stream.
 * 
 * Example:
 * 
 * 
 * int k = 3;
 * int[] arr = [4,5,8,2];
 * KthLargest kthLargest = new KthLargest(3, arr);
 * kthLargest.add(3);   // returns 4
 * kthLargest.add(5);   // returns 5
 * kthLargest.add(10);  // returns 5
 * kthLargest.add(9);   // returns 8
 * kthLargest.add(4);   // returns 8
 * 
 * 
 * Note: 
 * You may assume that nums' length ≥ k-1 and k ≥ 1.
 * 
 */
// Use k-size 'min-heap' to track the kth largest elements.
// When inserting an element, check if it has exceeds
// the capcity of min-heap:
//  * If not, just insert it into min-heap.
//  * If so, check if the element is larger than the root
//    of min-heap (O(1)), which is k-th large minimum element so far:
//      * If element is larger than min-heap's root, remove the
//        root element from min-heap and insert the element into it.
//      * If element is smaller or equal to min-heap's root,
//        then it must not be the k-th large number among the stream,
//        just discard it.
// Continue the iterations, then after all numbers in the stream
// are iterated, min-heap's root will be the k-th large number.
//
// Time Complexity of inserting element into heap: O(k), total n numbers,
// overall time complexity: O(nlogk)
// Space Complexity: O(k)
type KthLargest struct {
    minHeap *MinHeap
    k int
}

func Constructor(k int, nums []int) KthLargest {
    // Insert at most k elements into min-heap.
    n := min(k, len(nums))
    h := make(MinHeap, n)
    for i := 0; i < n; i++ {
        h[i] = nums[i]
    }

    this := KthLargest{
        minHeap: &h,
        k: k,
    }

    heap.Init(this.minHeap)

    // Insert the rest of elements.
    for i := n; i < len(nums); i++ {
        this.Add(nums[i])
    }

    return this
}

func (this *KthLargest) Add(val int) int {
    if len(*this.minHeap) == this.k {
        // Min-heap is full, check root.
        smallest := (*this.minHeap)[0]
        if val <= smallest {
            return (*this.minHeap)[0]
        }

        // Element is larger than min-heap's root, remove the
        // root element from min-heap and insert the element into it.
        (*this.minHeap)[0] = val
        heap.Fix(this.minHeap, 0)
    } else {
        // Min-heap is not full yet, just insert the element into min-heap.
        heap.Push(this.minHeap, val)
    }

    return (*this.minHeap)[0]
}

type MinHeap []int

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
