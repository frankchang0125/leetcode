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
// References:
//	http://bit.ly/2VOgIUI
//	http://bit.ly/2Xzn7Uv
//	http://bit.ly/2XtNOcS
type MyCalendar struct {
	root *Event
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

type Event struct {
	StartTime int
	EndTime   int
	Left      *Event
	Right     *Event
}

func (this *MyCalendar) InsertEvent(event *Event) {
	if this.root == nil {
		this.root = event
		return
	}

	current := this.root

	for {
		if event.StartTime < current.StartTime {
			if current.Left != nil {
				current = current.Left
			} else {
				current.Left = event
				return
			}
		} else {
			if current.Right != nil {
				current = current.Right
			} else {
				current.Right = event
				return
			}
		}
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	var event *Event

	event = this.floor(start)
	if event != nil && event.EndTime > start {
		return false
	}

	event = this.ceil(start)
	if event != nil && event.StartTime < end {
		return false
	}

	this.InsertEvent(&Event{
		StartTime: start,
		EndTime:   end,
	})

	return true
}

func (this *MyCalendar) floor(queryStartTime int) *Event {
	var floor *Event
	current := this.root

	for {
		if current == nil {
			break
		}

		if current.StartTime == queryStartTime {
			return current
		} else if current.StartTime > queryStartTime {
			current = current.Left
		} else {
			floor = current
			current = current.Right
		}
	}

	return floor
}

func (this *MyCalendar) ceil(queryStartTime int) *Event {
	var ceil *Event
	current := this.root

	for {
		if current == nil {
			break
		}

		if current.StartTime == queryStartTime {
			return current
		} else if current.StartTime > queryStartTime {
			ceil = current
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return ceil
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
