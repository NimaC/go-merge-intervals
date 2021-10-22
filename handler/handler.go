package handler

import (
	"sort"

	"github.com/NimaC/go-merge-intervals/interval"
	"github.com/NimaC/go-merge-intervals/stack"
)

func GetOverlappingIntervals(intervals []interval.Interval) []interval.Interval {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	intervalStack := stack.Stack{}
	intervalStack = intervalStack.Push(intervals[0])
	var stackElement interval.Interval
	for i := 1; i < len(intervals); i++ {
		intervalStack, stackElement, _ = intervalStack.Pop()
		if stackElement.End >= intervals[i].Start {
			mergedInterval := mergeIntervals(stackElement, intervals[i])
			intervalStack = intervalStack.Push(mergedInterval)
		} else {
			intervalStack = intervalStack.Push(stackElement)
			intervalStack = intervalStack.Push(intervals[i])
		}
	}
	return intervalStack
}

func mergeIntervals(s interval.Interval, e interval.Interval) interval.Interval {
	var elementEnd int
	if s.End >= e.End {
		elementEnd = s.End
	} else {
		elementEnd = e.End
	}
	return interval.Interval{s.Start, elementEnd}
}
