/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 *
 * https://leetcode.com/problems/sort-characters-by-frequency/description/
 *
 * algorithms
 * Medium (54.55%)
 * Total Accepted:    82.5K
 * Total Submissions: 150.4K
 * Testcase Example:  '"tree"'
 *
 * Given a string, sort it in decreasing order based on the frequency of
 * characters.
 *
 * Example 1:
 *
 * Input:
 * "tree"
 *
 * Output:
 * "eert"
 *
 * Explanation:
 * 'e' appears twice while 'r' and 't' both appear once.
 * So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid
 * answer.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * "cccaaa"
 *
 * Output:
 * "cccaaa"
 *
 * Explanation:
 * Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
 * Note that "cacaca" is incorrect, as the same characters must be together.
 *
 *
 *
 * Example 3:
 *
 * Input:
 * "Aabb"
 *
 * Output:
 * "bbAa"
 *
 * Explanation:
 * "bbaA" is also a valid answer, but "Aabb" is incorrect.
 * Note that 'A' and 'a' are treated as two different characters.
 *
 *
 */
type Freq struct {
	c    rune
	freq int
}

// Define 'FreqSlice' type and implement Sort interface
// to sort 'Freq' by 'freq'.
type FreqSlice []Freq

func (f FreqSlice) Len() int {
	return len(f)
}

func (f FreqSlice) Less(i, j int) bool {
	return f[i].freq < f[j].freq
}

func (f FreqSlice) Swap(i, j int) {
	tmp := f[i]
	f[i] = f[j]
	f[j] = tmp
}

// Reference: https://goo.gl/jorr5o
func frequencySort(s string) string {
	freq := make(FreqSlice, 128)

	// Count the frequency of each character in the string.
	for _, c := range s {
		freq[c].c = c
		freq[c].freq++
	}

	// Sort the characters in the string by frequency
	// from high to low.
	sort.Sort(sort.Reverse(freq))

	var buf bytes.Buffer

	// Output the sorted characters.
	for _, f := range freq {
		if f.freq == 0 {
			break
		}

		for i := 0; i < f.freq; i++ {
			buf.WriteByte(byte(f.c))
		}
	}

	return buf.String()
}
