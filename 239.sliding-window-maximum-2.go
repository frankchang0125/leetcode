/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 *
 * https://leetcode.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (38.02%)
 * Likes:    1652
 * Dislikes: 95
 * Total Accepted:    157.5K
 * Total Submissions: 413.2K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * Given an array nums, there is a sliding window of size k which is moving
 * from the very left of the array to the very right. You can only see the k
 * numbers in the window. Each time the sliding window moves right by one
 * position. Return the max sliding window.
 * 
 * Example:
 * 
 * 
 * Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
 * Output: [3,3,5,5,6,7] 
 * Explanation: 
 * 
 * Window position                Max
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 * 
 * 
 * Note: 
 * You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty
 * array.
 * 
 * Follow up:
 * Could you solve it in linear time?
 */
// To find out the maximum number in sliding window, we can use 'queue' to keep
// the 'index' of numbers.
// (The reason why we have to keep the 'index' of numbers instead of the
//  'value' of numbers is because there might be duplicate numbers in array.
//  If we just keep the values of numbers, we are not able to determine whether
//  the number at the queue's front is the slide out number or not.)
//
// We have to make sure the maximum number's index in sliding window is always at
// the front of queue. In order to ensure this:
//  * When inserting a new number's index:
//      => We have to make sure there're no numbers which is smaller than
//         the inserted number. If there're any numbers which are smaller than
//         the inserted number, since they won't be the maximum values in
//         sliding window, pop all of them from queue's back.
//  * When the maximum number's index (i.e. Queue's front) is going to slide out
//    from sliding window:
//      => We have to remove it from queue, too.
//
//  ex: nums: [2, 3, 4, 2, 6, 2, 5, 1], sliding window size: 3
//      (Note: Queue here shows the value of number instead of index.)
//
//                                                      *
//      +-------+-----------------+----------------+---------+---------------+
//      | Steps | Inserted nubmer | Sliding window |  Queue  | Maximum value |
//      +-------+-----------------+----------------+---------+---------------+
//      |   1   |        2        | [2]            | [2]     |      N/A      |
//      |   2   |        3        | [2, 3]         | [3]     |      N/A      |
//      |   3   |        4        | [2, 3, 4]      | [4]     |       4       |
//      |   4   |        2        | [3, 4, 2]      | [4, 2]  |       4       |
//      |   5   |        6        | [4, 2, 6]      | [6]     |       6       |
//      |   6   |        2        | [2, 6, 2]      | [6, 2]  |       6       |
//      |   7   |        5        | [6, 2, 5]      | [6, 5]  |       6       |
//      |   8   |        1        | [2, 5, 1]      | [5, 1]  |       5       |
//      +-------+-----------------+----------------+---------+---------------+
//
// Time Complexity: O(n)
// Space Complexity: O(k)
func maxSlidingWindow(nums []int, k int) []int {
    maxWindows := []int{}

    if len(nums) >= k && k >= 1 {
        // Dequeue which keeps the 'index' of numbers.
        deque := list.New()

        // Fill sliding window first.
        // The maximum number's index in the sliding window will
        // at the front of queue.
        for i := 0; i < k; i++ {
            // Pop all numbers from queue's back which are smaller than
            // the inserted number.
            for deque.Len() > 0 && nums[i] > nums[deque.Back().Value.(int)] {
                deque.Remove(deque.Back())
            }
            // Push the inserted number to queue's back.
            deque.PushBack(i)
        }

        for i := k; i < len(nums); i++ {
            maxWindows = append(maxWindows, nums[deque.Front().Value.(int)])

            // Pop all numbers from queue's back which are smaller than
            // the inserted number.
            for deque.Len() > 0 && nums[i] >= nums[deque.Back().Value.(int)] {
                deque.Remove(deque.Back())
            }

            // If the maximum number (i.e. Queue's front) is going to slide out
            // from sliding window, remove it from queue.
            if deque.Len() > 0 && deque.Front().Value.(int) <= i - k {
                deque.Remove(deque.Front())
            }

            // Push the inserted number to queue's back.
            deque.PushBack(i)
        }

        // Add the last maximum number to answer.
        maxWindows = append(maxWindows, nums[deque.Front().Value.(int)])
    }

    return maxWindows
}
