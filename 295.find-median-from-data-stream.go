/*
 * @lc app=leetcode id=295 lang=golang
 *
 * [295] Find Median from Data Stream
 *
 * https://leetcode.com/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (36.35%)
 * Likes:    1113
 * Dislikes: 23
 * Total Accepted:    106.9K
 * Total Submissions: 292.9K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n[[],[1],[2],[],[3],[]]'
 *
 * Median is the middle value in an ordered integer list. If the size of the
 * list is even, there is no middle value. So the median is the mean of the two
 * middle value.
 * For example,
 * 
 * [2,3,4], the median is 3
 * 
 * [2,3], the median is (2 + 3) / 2 = 2.5
 * 
 * Design a data structure that supports the following two operations:
 * 
 * 
 * void addNum(int num) - Add a integer number from the data stream to the data
 * structure.
 * double findMedian() - Return the median of all elements so far.
 * 
 * 
 * 
 * 
 * Example:
 * 
 * 
 * addNum(1)
 * addNum(2)
 * findMedian() -> 1.5
 * addNum(3) 
 * findMedian() -> 2
 * 
 * 
 * 
 * 
 * Follow up:
 * 
 * 
 * If all integer numbers from the stream are between 0 and 100, how would you
 * optimize it?
 * If 99% of all integer numbers from the stream are between 0 and 100, how
 * would you optimize it?
 * 
 * 
 */
// To find the median from data stream,
// we can divide numbers into left and right parts:
//  Numbers smaller or equal to median.
//  Numbers larger or equal to median.
//
// For left part numbers, we can use 'max-heap' to store
// the numbers which are smaller or equal to median.
// For right part numbers, we can use 'min-heap' to store
// the numbers which are larger or equal to median.
//
// For odd-length data stream:
//  Size of left part max-heap must not 1-element larger or smaller
//  than right part min-heap.
//  The median will be the root element of:
//      Left part max-heap, if size of left part max-heap
//      is 1-element larger than right part min-heap.
//      Right part min-heap, if size of right part min-heap
//      is 1-element larger than left part max-heap.
//  
// For even-length data stream:
//  Size of left part max-heap must be same as right part min-heap.
//  The median will be the mean of left part max-heap root element
//  and right part min-heap root element.
//
// In this solution, we choose to let left-part max-heap to be 1-element
// larger than right-part min-heap for odd-length data stream.
type MedianFinder struct {
    left *MaxHeap
    right *MinHeap
    count int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
    maxHeap := make(MaxHeap, 0)
    minHeap := make(MinHeap, 0)

    return MedianFinder {
        left: &maxHeap,
        right: &minHeap,
        count: 0,
    }
}

func (this *MedianFinder) AddNum(num int)  {
    defer func() {
        this.count++
    }()

    if this.count == 0 {
        heap.Push(this.left, num)
        return
    }

    if this.count & 1 == 0 {
        // Even-length data stream.
        leftMax := this.left.Root().(int)

        if num <= leftMax {
            // Insert number is less or equal to left-part max-heap's root,
            // just insert the number into left-part max-heap.
            //  ex: [2, 3] [5, 6], insert num: 1
            //      => [1, 2, 3] [5, 6]
            //  ex: [2, 3] [5, 6], insert num: 3
            //      => [2, 3, 3] [5, 6]
            heap.Push(this.left, num)
        } else {
            // Insert number is larger than left-part max-heap's root,
            // rebalance the heaps by first insert the number into
            // right-part min-heap then move right-part min-heap's root
            // to left-part max-heap.
            //  ex: [2, 3] [5, 6], insert num: 4
            //      => [2, 3] [4, 5, 6]
            //      => [2, 3, 4] [5, 6]
            //  ex: [2, 3] [5, 6], insert num: 5
            //      => [2, 3] [5, 5, 6]
            //      => [2, 3, 5] [5, 6]
            //  ex: [2, 3] [5, 6], insert num: 7
            //      => [2, 3] [5, 6, 7]
            //      => [2, 3, 5] [6, 7]
            heap.Push(this.right, num)
            rightMin := heap.Remove(this.right, 0)
            heap.Push(this.left, rightMin)
        }
    } else {
        // Odd-length data stream.
        leftMax := this.left.Root().(int)

        if num >= leftMax {
            // Insert number is larger or equal to left-part max-heap's root,
            // just insert the number into right-part min-heap.
            //  ex: [2, 3] [5], insert num: 3
            //      => [2, 3] [3, 5]
            //  ex: [2, 3] [5], insert num: 4
            //      => [2, 3] [4, 5]
            heap.Push(this.right, num)
        } else {
            // Insert number is smaller than left-part max-heap's root,
            // rebalance the heaps by first insert the number into
            // left-part max-heap then move left-part max-heap's root
            // to right-part min-heap.
            //  ex: [2, 3] [5], insert num: 1
            //      => [1, 2, 3] [5]
            //      => [1, 2] [3, 5]
            heap.Push(this.right, leftMax)
            heap.Remove(this.left, 0)
            heap.Push(this.left, num)
        }
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.count == 0 {
        return -1
    }

    if this.count & 1 == 0 {
        // Even-length data stream.
        leftMax := this.left.Root().(int)
        rightMin := this.right.Root().(int)
        return (float64(leftMax) + float64(rightMin)) / 2
    } else {
        // Odd-length data stream.
        return float64(this.left.Root().(int))
    }
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

func (h MinHeap) Root() interface{} {
    if len(h) == 0 {
        return 0
    }
    return h[0]
}

type MaxHeap []int

func (h MaxHeap) Len() int {
    return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
    return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h MaxHeap) Root() interface{} {
    if len(h) == 0 {
        return 0
    }
    return h[0]
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
