/*
 * @lc app=leetcode id=506 lang=golang
 *
 * [506] Relative Ranks
 *
 * https://leetcode.com/problems/relative-ranks/description/
 *
 * algorithms
 * Easy (47.77%)
 * Total Accepted:    41.2K
 * Total Submissions: 85.7K
 * Testcase Example:  '[5,4,3,2,1]'
 *
 *
 * Given scores of N athletes, find their relative ranks and the people with
 * the top three highest scores, who will be awarded medals: "Gold Medal",
 * "Silver Medal" and "Bronze Medal".
 *
 * Example 1:
 *
 * Input: [5, 4, 3, 2, 1]
 * Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
 * Explanation: The first three athletes got the top three highest scores, so
 * they got "Gold Medal", "Silver Medal" and "Bronze Medal". For the left two
 * athletes, you just need to output their relative ranks according to their
 * scores.
 *
 *
 *
 * Note:
 *
 * N is a positive integer and won't exceed 10,000.
 * All the scores of athletes are guaranteed to be unique.
 *
 *
 *
 */
func findRelativeRanks(nums []int) []string {
	m := map[int]int{}
	ranks := make([]string, len(nums))

	for i, num := range nums {
		m[num] = i
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	for i, num := range nums {
		pos := m[num]

		switch i {
		case 0:
			ranks[pos] = "Gold Medal"
		case 1:
			ranks[pos] = "Silver Medal"
		case 2:
			ranks[pos] = "Bronze Medal"
		default:
			ranks[pos] = strconv.Itoa(i + 1)
		}
	}

	return ranks
}
