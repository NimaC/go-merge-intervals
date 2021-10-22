package stack

import (
	"../interval"
)

type stack []interval.Interval

func (s stack) Push(v interval.Interval) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, interval.Interval, bool) {
	l := len(s)
	if l == 0 {
		return s, interval.Interval{0, 0}, false
	}
	return s[:l-1], s[l-1], true
}
