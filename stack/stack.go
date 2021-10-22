package stack

import (
	"github.com/NimaC/go-merge-intervals/interval"
)

// Minimal Stack Implementation for Interval Type with Push and Pop functionality
type Stack []interval.Interval

func (s Stack) Push(v interval.Interval) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, interval.Interval, bool) {
	l := len(s)
	if l == 0 {
		return s, interval.Interval{0, 0}, false
	}
	return s[:l-1], s[l-1], true
}
