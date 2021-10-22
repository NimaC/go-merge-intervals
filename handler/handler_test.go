package handler

import (
	"reflect"
	"testing"

	"github.com/NimaC/go-merge-intervals/interval"
)

func generalTest(t *testing.T, input []interval.Interval, expectedOutput []interval.Interval) {
	overlappingIntervals := GetOverlappingIntervals(input)
	if !reflect.DeepEqual(overlappingIntervals, expectedOutput) {
		t.Errorf("Test Failed. Expected: %d Output: %d", expectedOutput, overlappingIntervals)
	}
}

// Input: []  Output: []
func TestEmpty(t *testing.T) {
	var input []interval.Interval
	generalTest(t, input, input)
}

// Input: [0, 1]  Output: [0, 1]
func TestSizeOne(t *testing.T) {
	input := []interval.Interval{
		{0, 1},
	}
	generalTest(t, input, input)
}

// Input: [25,30] [2,19] [14, 23] [4,8]  Output: [2,23] [25,30]
func TestBriefingExample(t *testing.T) {
	intervals := []interval.Interval{
		{25, 30},
		{2, 19},
		{14, 23},
		{4, 8},
	}
	expectedOutput := []interval.Interval{
		{2, 23},
		{25, 30},
	}
	generalTest(t, intervals, expectedOutput)
}

func TestNoOverlaps(t *testing.T) {
	intervals := []interval.Interval{
		{0, 1},
		{4, 8},
		{10, 34},
		{55, 60},
	}
	expectedOutput := intervals
	generalTest(t, intervals, expectedOutput)
}

func TestBiggerExample(t *testing.T) {
	intervals := []interval.Interval{
		{2, 2},
		{0, 2},
		{15, 23},
		{7, 15},
		{5, 7},
		{2, 3},
	}
	expectedOutput := []interval.Interval{
		{0, 3},
		{5, 23},
	}
	generalTest(t, intervals, expectedOutput)
}
