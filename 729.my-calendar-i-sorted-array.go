/*
 * @lc app=leetcode id=729 lang=golang
 *
 * [729] My Calendar I
 *
 * https://leetcode.com/problems/my-calendar-i/description/
 *
 * algorithms
 * Medium (45.97%)
 * Total Accepted:    29.5K
 * Total Submissions: 62.9K
 * Testcase Example:  '["MyCalendar","book","book","book"]\n[[],[10,20],[15,25],[20,30]]'
 *
 * Implement a MyCalendar class to store your events. A new event can be added
 * if adding the event will not cause a double booking.
 *
 * Your class will have the method, book(int start, int end). Formally, this
 * represents a booking on the half open interval [start, end), the range of
 * real numbers x such that start <= x < end.
 *
 * A double booking happens when two events have some non-empty intersection
 * (ie., there is some time that is common to both events.)
 *
 * For each call to the method MyCalendar.book, return true if the event can be
 * added to the calendar successfully without causing a double booking.
 * Otherwise, return false and do not add the event to the calendar.
 * Your class will be called like this: MyCalendar cal = new MyCalendar();
 * MyCalendar.book(start, end)
 *
 * Example 1:
 *
 *
 * MyCalendar();
 * MyCalendar.book(10, 20); // returns true
 * MyCalendar.book(15, 25); // returns false
 * MyCalendar.book(20, 30); // returns true
 * Explanation:
 * The first event can be booked.  The second can't because time 15 is already
 * booked by another event.
 * The third event can be booked, as the first event takes every time less than
 * 20, but not including 20.
 *
 *
 *
 *
 * Note:
 *
 *
 * The number of calls to MyCalendar.book per test case will be at most
 * 1000.
 * In calls to MyCalendar.book(start, end), start and end are integers in the
 * range [0, 10^9].
 *
 *
 *
 *
 */
// Overlapped events:
//  1.  -------------
//     | s1       e1 |
//      -------------
//          --------------
//         | s2        e2 |
//          --------------
//
//  2.  ------------------
//     | s1            e1 |
//      ------------------
//         -----------
//        | s2     e2 |
//         -----------
//
//  3.      -------------
//         | s1       e1 |
//          -------------
//      -------------
//     | s2       e2 |
//      -------------
//
//  4.     -----------
//        | s1     e1 |
//         -----------
//      ------------------
//     | s2            e2 |
//      ------------------
//
// floor(query): Event with 'largest' StartTime, where event.StartTime <= query.StartTime.
// ceil(query): Event with 'smallest' StartTime, where event.StartTime >= query.StartTime.
//
// For case 1, 2:
//  If floor(query.StartTime).EndTime > query.StartTime, they are overlapped.
//  i.e. If floor(query.StartTime).EndTime <= query.StartTime, they are not overlapped.
//
// For case 3, 4:
//  If ceil(query.StartTime).StartTime < query.EndTime, they are overlapped.
//  i.e. If cell(query.StartTime).StartTime >= query.EndTime, they are not overlapped.
//
// Use map to store the relation of StartTime and EndTime for each event.
//
// References:
//	http://bit.ly/2VOgIUI
//	http://bit.ly/2Xzn7Uv
//	http://bit.ly/2XtNOcS
type MyCalendar struct {
	// key: StartTime, value: EndTime.
	Events map[int]int
	// No sorted map in Go, use array for sorting.
	StartTimes []int
}

func Constructor() MyCalendar {
	return MyCalendar{
		Events:     map[int]int{},
		StartTimes: []int{},
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	var key int

	key = this.floor(start)
	if key != -1 && this.Events[key] > start {
		return false
	}

	key = this.ceil(start)
	if key != -1 && key < end {
		return false
	}

	this.Events[start] = end
	this.StartTimes = append(this.StartTimes, start)

	// Array must be sorted for binary search.
	sort.Sort(sort.IntSlice(this.StartTimes))

	return true
}

func (this *MyCalendar) floor(queryStartTime int) int {
	l := 0
	r := len(this.StartTimes)
	floor := -1

	for {
		if l >= r {
			break
		}

		m := (l + r) / 2

		if this.StartTimes[m] == queryStartTime {
			return this.StartTimes[m]
		} else if this.StartTimes[m] > queryStartTime {
			r = m
		} else {
			floor = this.StartTimes[m]
			l = m + 1
		}
	}

	return floor
}

func (this *MyCalendar) ceil(queryStartTime int) int {
	l := 0
	r := len(this.StartTimes)
	ceil := -1

	for {
		if l >= r {
			break
		}

		m := (l + r) / 2

		if this.StartTimes[m] == queryStartTime {
			return this.StartTimes[m]
		} else if this.StartTimes[m] > queryStartTime {
			ceil = this.StartTimes[m]
			r = m
		} else {
			l = m + 1
		}
	}

	return ceil
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
