/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 *
 * https://leetcode.com/problems/partition-labels/description/
 *
 * algorithms
 * Medium (67.83%)
 * Total Accepted:    35.1K
 * Total Submissions: 51.6K
 * Testcase Example:  '"ababcbacadefegdehijhklij"'
 *
 *
 * A string S of lowercase letters is given.  We want to partition this string
 * into as many parts as possible so that each letter appears in at most one
 * part, and return a list of integers representing the size of these parts.
 *
 *
 * Example 1:
 *
 * Input: S = "ababcbacadefegdehijhklij"
 * Output: [9,7,8]
 * Explanation:
 * The partition is "ababcbaca", "defegde", "hijhklij".
 * This is a partition so that each letter appears in at most one part.
 * A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it
 * splits S into less parts.
 *
 *
 *
 * Note:
 * S will have length in range [1, 500].
 * S will consist of lowercase letters ('a' to 'z') only.
 *
 */
// Reference: https://goo.gl/CMT2Qr
func partitionLabels(S string) []int {
	// Preprocessing, go through the string and find out the last index of
	// each character appeared in S.
	lastIndex := make([]int, 26)
	for i, c := range S {
		lastIndex[c-'a'] = i
	}

	result := []int{}
	start := 0 // The start index of current partition.
	end := 0   // The end index of current partition.

	for i := 0; i < len(S); i++ {
		// If current character's last index is larger than current partition's
		// end boundary, we have to extend the size of partition by updating
		// end's value.
		last := lastIndex[S[i]-'a']
		if last > end {
			end = last
		}

		// We have reached the end of partition, add the length of partition
		// to the result and start iterating next partition.
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}

	return result
}
