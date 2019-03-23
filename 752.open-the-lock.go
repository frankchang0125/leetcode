/*
 * @lc app=leetcode id=752 lang=golang
 *
 * [752] Open the Lock
 *
 * https://leetcode.com/problems/open-the-lock/description/
 *
 * algorithms
 * Medium (43.82%)
 * Total Accepted:    17.7K
 * Total Submissions: 39.5K
 * Testcase Example:  '["0201","0101","0102","1212","2002"]\n"0202"'
 *
 *
 * You have a lock in front of you with 4 circular wheels.  Each wheel has 10
 * slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'.  The wheels can
 * rotate freely and wrap around: for example we can turn '9' to be '0', or '0'
 * to be '9'.  Each move consists of turning one wheel one slot.
 *
 * The lock initially starts at '0000', a string representing the state of the
 * 4 wheels.
 *
 * You are given a list of deadends dead ends, meaning if the lock displays any
 * of these codes, the wheels of the lock will stop turning and you will be
 * unable to open it.
 *
 * Given a target representing the value of the wheels that will unlock the
 * lock, return the minimum total number of turns required to open the lock, or
 * -1 if it is impossible.
 *
 *
 * Example 1:
 *
 * Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
 * Output: 6
 * Explanation:
 * A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" ->
 * "1201" -> "1202" -> "0202".
 * Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202"
 * would be invalid,
 * because the wheels of the lock become stuck after the display becomes the
 * dead end "0102".
 *
 *
 *
 * Example 2:
 *
 * Input: deadends = ["8888"], target = "0009"
 * Output: 1
 * Explanation:
 * We can turn the last wheel in reverse to move from "0000" -> "0009".
 *
 *
 *
 * Example 3:
 *
 * Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"],
 * target = "8888"
 * Output: -1
 * Explanation:
 * We can't reach the target without getting stuck.
 *
 *
 *
 * Example 4:
 *
 * Input: deadends = ["0000"], target = "8888"
 * Output: -1
 *
 *
 *
 * Note:
 *
 * The length of deadends will be in the range [1, 500].
 * target will not be in the list deadends.
 * Every string in deadends and the string target will be a string of 4 digits
 * from the 10,000 possibilities '0000' to '9999'.
 *
 *
 */
// Use BFS to find a sequence of valid moves to open the lock.
// Reference: https://goo.gl/rh6V8r
func openLock(deadends []string, target string) int {
	deadendsMap := make(map[string]bool)
	for _, deadend := range deadends {
		if deadend == "0000" {
			return -1
		}
		deadendsMap[deadend] = true
	}

	visited := make(map[string]bool)

	queue := []string{"0000"} // Enqueue
	steps := 0

	for {
		if len(queue) == 0 {
			break
		}

		steps++

		for i := len(queue); i > 0; i-- {
			num := queue[0] // Dequeue
			queue = queue[1:]

			var n []byte
			var nStr string

			// Turn up and down for each wheel.
			for i := 0; i < 4; i++ {
				n = []byte(num)

				// Turn the wheel up.
				if n[i] == '9' {
					n[i] = '0'
				} else {
					n[i] = n[i] + 1
				}

				nStr = string(n)
				if nStr == target {
					// Found the target, return steps.
					return steps
				}

				_, dead := deadendsMap[nStr]
				_, visit := visited[nStr]
				if !dead && !visit {
					// Enqueue the new code for next BFS
					// if the code is not deadend and not visited before.
					queue = append(queue, nStr) // Enqueue
					visited[nStr] = true
				}

				// Turn the wheel down.
				if n[i] == '1' {
					n[i] = '9'
				} else if n[i] == '0' {
					n[i] = '8'
				} else {
					n[i] = n[i] - 2
				}

				nStr = string(n)
				if nStr == target {
					// Found the target, return steps.
					return steps
				}

				_, dead = deadendsMap[nStr]
				_, visit = visited[nStr]
				if !dead && !visit {
					// Enqueue the new code for next BFS
					// if the code is not deadend and not visited before.
					queue = append(queue, nStr) // Enqueue
					visited[nStr] = true
				}
			}
		}
	}

	// Cannot find valid sequence to open the lock.
	return -1
}
